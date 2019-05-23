//#include <stdio.h>
//#include <stdlib.h>
//#include <unistd.h>
////#include <libtask.h>
//
//int quiet, goal, buffer;
//
//void primetask(void *arg)
//{
//    Channel *c, *nc;
//    int p, i;
//    c = arg;
//
//    p = chanrecvul(c);
//    if (p > goal)
//        taskexitall(0);
//    if (!quiet)
//        printf("%d\n", p);
//    nc = chancreate(sizeof(unsigned long), buffer);
//    taskcreate(primetask, nc, 32768)
//    for (;;) {
//        i = chanrecvul(c);
//        if (i%p) {
//            chansendul(nc, i);
//        }
//    }
//
//}