#ifndef SHM_H
#define SHM_H

#include <stdlib.h>
#include <sys/shm.h>
#include <sys/types.h>
#include <sys/ipc.h>

int sysv_shm_open();
size_t sysv_shm_get_size(int shm_id);
void *sysv_shm_attach(int shm_id);

// SHM_H
#endif
