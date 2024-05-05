#include <liburing.h>
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>

#define QUEUE_SIZE 10
#define CHUNK_BYTE_SIZE 1024  // 1KB
#define PATH "example.txt"

enum Stage {
  STAGE_START,
  STAGE_FILE_OPEN,
  STAGE_QUEUE_INITIALIZED,
  STAGE_CHUNKS_ALLOCATED
};

int main() {
  int exit_code = 0;
  enum Stage stage = STAGE_START;

  int fd = open(PATH, O_RDONLY);
  if (fd == -1) {
    fprintf(stderr, "failed to open file at %s: %s\n", PATH, strerror(errno));
    exit_code = 1;
    goto CLEANUP;
  }
  stage = STAGE_FILE_OPEN;

  struct stat file_info;
  if (fstat(fd, &file_info) < 0) {
    fprintf(stderr, "failed to get file info: %s\n", strerror(errno));
    exit_code = 1;
    goto CLEANUP;
  }
  off_t size = file_info.st_size;
  uint n_chunks = (uint)ceilf((float)size / CHUNK_BYTE_SIZE);
  printf("File info (size: %ld, chunk_size: %d, n_chunks: %d)\n", size, CHUNK_BYTE_SIZE, n_chunks);

  printf("Allocating %d chunks\n", n_chunks);
  char **chunks = malloc(sizeof(char*) * n_chunks);
  if (chunks == NULL) {
    fprintf(stderr, "malloc: %s\n", strerror(errno));
    exit_code = 1;
    goto CLEANUP;
  }
  stage = STAGE_CHUNKS_ALLOCATED;

  printf("Initializing ring queue\n");
  struct io_uring ring;
  int status = io_uring_queue_init(QUEUE_SIZE, &ring, 0);
  if (status < 0) {
    fprintf(stderr, "queue_init: %s\n", strerror(-status));
    exit_code = 1;
    goto CLEANUP;
  }
  stage = STAGE_QUEUE_INITIALIZED;

  printf("Creating an SQE for each chunk\n");
  struct io_uring_sqe *sqe;
  off_t offset = 0;
  for (uint i = 0; i < n_chunks; i++) {
    // Get the next available submission queue entry
    sqe = io_uring_get_sqe(&ring);
    if (sqe == NULL) {
      fprintf(stderr, "ring queue is full\n");
      exit_code = 1;
      goto CLEANUP;
    }

    // Alloc chunk memory
    chunks[i] = malloc(CHUNK_BYTE_SIZE);

    // Read file chunck
    io_uring_prep_read(sqe, fd, chunks[i], CHUNK_BYTE_SIZE, offset);
    offset += CHUNK_BYTE_SIZE;
  }

  printf("Submitting queue\n");
  int submitted_requests = io_uring_submit(&ring);
  if (submitted_requests < 0) {
    fprintf(stderr, "io_uring_submit: %s\n", strerror(-submitted_requests));
    exit_code = 1;
    goto CLEANUP;
  } else if (submitted_requests != n_chunks) {
    fprintf(
      stderr,
      "io_uring_submit: created %d requests but %d were submitted\n",
      n_chunks, submitted_requests
    );
    exit_code = 1;
    goto CLEANUP;
  }

  printf("Wait for CQEs and marking them as seen\n");
  struct io_uring_cqe *cqe;
  offset = 0;
  for (int i = 0; i < n_chunks; i++) {
    status = io_uring_wait_cqe(&ring, &cqe);
    if (status < 0) {
      fprintf(stderr, "io_uring_wait_cqe: %s\n", strerror(-status));
      exit_code = 1;
      goto CLEANUP;
    }

    if (cqe->res != CHUNK_BYTE_SIZE && offset + cqe->res != file_info.st_size) {
      fprintf(
        stderr,
        "expected to read %d bytes but read %d bytes (file was not fully read)\n",
        CHUNK_BYTE_SIZE, cqe->res
      );
      exit_code = 1;
      goto CLEANUP;
    }
    offset += cqe->res;
    io_uring_cqe_seen(&ring, cqe);

    printf("---------- CHUNK %d ----------\n", i+1);
    puts(chunks[i]);
  }

CLEANUP:
  if (stage >= STAGE_QUEUE_INITIALIZED) io_uring_queue_exit(&ring);
  if (stage >= STAGE_CHUNKS_ALLOCATED) {
    for (uint i = 0; i < n_chunks; i++) {
      if (chunks[i] != NULL) free(&chunks[i]);
    }
    free(chunks);
  }
  if (stage >= STAGE_FILE_OPEN) close(fd);

  return exit_code;
}