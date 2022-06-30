#include "shm.h"

int sysv_shm_open() {
    int shm_id;

    key_t key = ftok("/tmp", 2333);

    if (key < 0) {
        return -1;
    }

    return shmget(key, 0, 0);
}

size_t sysv_shm_get_size(int shm_id) {
    struct shmid_ds shm;

    if (shmctl(shm_id, IPC_STAT, &shm) >= 0) {
        return shm.shm_segsz;
    }
    else {
        return -1;
    }
}

void *sysv_shm_attach(int shm_id) {
    return shmat(shm_id, NULL, 0);
}


