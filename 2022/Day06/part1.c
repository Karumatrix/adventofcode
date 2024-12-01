#include <fcntl.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>
#include <sys/stat.h>
#include <stddef.h>
#include <string.h>

char *load_file_in_mem(void)
{
    struct stat s;
    stat("./input.txt", &s);
    char *content = malloc(sizeof(char) * s.st_size);
    int fd = open("./input.txt", O_RDONLY);
    if (read(fd, content, s.st_size) <= 0) {
        free(content);
        close(fd);
        return NULL;
    }
    close(fd);
    return content;
}

int is_same_char(int len, int i, char const *filecontent)
{
    for (int j = 0; j < len - 1; j++) {
        for (int k = j + 1; k < len; k++) {
            if (filecontent[i + j] == filecontent[i + k]) {
                return 0;
            }
        }
    }
    return i + len;
}

int main(void)
{
    int len = 14;
    int result = 0;
    char *filecontent = load_file_in_mem();
    for (int i = 0; filecontent[i + len] != '\0'; i++) {
        if (is_same_char(len, i, filecontent) != 0) {
            result = is_same_char(len, i, filecontent);
            break;
        }
    }
    printf("%d\n", result);
    free(filecontent);
    return 0;
}