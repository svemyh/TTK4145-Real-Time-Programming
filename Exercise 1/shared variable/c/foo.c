// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t mutex;   //Declare a mutex object

// Note the return type: void*
void* incrementingThreadFunction(){
    // TODO: increment i 1_000_000 times
    pthread_mutex_lock(&mutex);
    for (int j = 0; j < 1000000; j++) {
        i++;
    }
    pthread_mutex_unlock(&mutex);
    return NULL;
}

void* decrementingThreadFunction(){
    // TODO: decrement i 1_000_000 times
    pthread_mutex_lock(&mutex);
    for (int j = 0; j < 1000000; j++) {
        i--;
    }
    pthread_mutex_unlock(&mutex);
    return NULL;
}

/**
 * Ref: https://www.geeksforgeeks.org/thread-functions-in-c-c/

int pthread_create(pthread_t * thread, 
            const pthread_attr_t * attr, 
            void * (*start_routine)(void *), 
            void *arg);

Parameters:

    thread: pointer to an unsigned integer value that returns the thread id of the thread created.
    attr: pointer to a structure that is used to define thread attributes like detached state, scheduling policy, stack address, etc. Set to NULL for default thread attributes.
    start_routine: pointer to a subroutine that is executed by the thread. The return type and parameter type of the subroutine must be of type void *. The function has a single attribute but if multiple values need to be passed to the function, a struct must be used.
    arg: pointer to void that contains the arguments to the function defined in the earlier argument

int pthread_join(pthread_t th, 
                 void **thread_return);

Parameter: This method accepts following parameters:

    th: thread id of the thread for which the current thread waits.
    thread_return: pointer to the location where the exit status of the thread mentioned in th is stored.


 * */



int main(){
    // TODO: 
    // start the two functions as their own threads using `pthread_create`
    // Hint: search the web! Maybe try "pthread_create example"?
    pthread_t incrementingThread, decrementingThread; // initialize the threads
    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL); // Using thread with no special attributes, and "coupling it" with a function with no special attributes.
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);
    
    // TODO:
    // wait for the two threads to be done before printing the final result
    // Hint: Use `pthread_join`    
    // Waiting for threads to finish
    pthread_join(incrementingThread, NULL);
    pthread_join(decrementingThread, NULL);
    


    printf("The magic number is: %d\n", i);
    return 0;
}
