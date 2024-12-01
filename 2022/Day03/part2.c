#include <fcntl.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>
#include <sys/stat.h>
#include <stddef.h>
#include <string.h>

char **load_2d_arr_from_file(char *filecontent)
{
    int num_line = 0;
    int num_char = 0;
    int index = 0;
    int temp_index = 0;
    int result_index = 0;
    int temp_result_index = 0;
    while (filecontent[index] != '\0') {
        if (filecontent[index] == '\n')
            num_line += 1;
        index += 1;
    }
    char **result = malloc(sizeof(char *) * (num_line + 2));
    index = 0;
    while (filecontent[index] != '\0') {
        temp_index = index;
        num_char = 0;
        while (filecontent[temp_index] != '\0') {
            if (filecontent[temp_index] == '\n')
                break;
            temp_index++;
            num_char++;
        }
        result[result_index] = malloc(sizeof(char) * (num_char + 1));
        result[result_index][0] = '\0';
        temp_result_index = 0;
        while (filecontent[index] != '\0') {
            if (filecontent[index] == '\n') {
                index++;
                break;
            }
            result[result_index][temp_result_index] = filecontent[index];
            temp_result_index++;
            index++;
        }
        result[result_index][num_char] = '\0';
        result_index++;
    }
    result[num_line + 1] = NULL;
    free(filecontent);
    return result;
}

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

char find_same_item(char const *line1, char const *line2, char const *line3)
{
    for (int i = 0; line1[i] != '\0'; i++) {
        for (int j = 0; line2[j] != '\0'; j++) {
            if (line1[i] == line2[j]) {
                for (int k = 0; line3[k] != '\0'; k++) {
                    if (line2[j] == line3[k])
                        return line3[k];
                }
            }
        }
    }
}

int main(void)
{
    int result = 0;
    char **filecontent = load_2d_arr_from_file(load_file_in_mem());
    for (int i = 0; filecontent[i] != NULL; i += 3) {
        char priority = find_same_item(filecontent[i], filecontent[i + 1], filecontent[i + 2]);
        if (priority >= 'A' && priority <= 'Z') {
            result += priority - 'A' + 27;
        }
        if (priority >= 'a' && priority <= 'z') {
            result += priority - 'a' + 1;
        }
    }
    printf("%d\n", result);
    for (int i = 0; filecontent[i] != NULL; i++)
        free(filecontent[i]);
    free(filecontent);
    return 0;
}