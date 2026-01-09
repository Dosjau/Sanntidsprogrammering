// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t lock = PTHREAD_MUTEX_INITIALIZER;

// Note the return type: void*
void* incrementingThreadFunction(){
    for (int j = 0; j < 1000001; j++){
        pthread_mutex_lock(&lock);
        i++;
        pthread_mutex_unlock(&lock);
    }
    return NULL;
}

void* decrementingThreadFunction(){
    for (int j = 0; j < 1000000; j++){
        pthread_mutex_lock(&lock);
        i--;
        pthread_mutex_unlock(&lock);
    }
    return NULL;
}


int main(){
    
    pthread_t inc_thread, dec_thread;

    pthread_create(&inc_thread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&dec_thread, NULL, decrementingThreadFunction, NULL);

    pthread_join(inc_thread, NULL);
    pthread_join(dec_thread, NULL);
   
    printf("The magic number is: %d\n", i);
    return 0;
}
