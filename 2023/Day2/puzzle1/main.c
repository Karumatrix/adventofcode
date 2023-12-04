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

int find_substring(char *s_string, char *r_string, int r_index)
{
    for (int j = 0; s_string[j] != '\0'; j++) {
        if (r_string[j + r_index] == '\0')
            return 0;
        if (s_string[j] != r_string[j + r_index])
            return 0;
    }
    return 1;
}

int main(void)
{
    int result = 0;
    int max_red = 12;
    int max_green = 13;
    int max_blue = 14;
    char **filecontent = load_2d_arr_from_file(load_file_in_mem());
    for (int i = 0; filecontent[i] != NULL; i++) {
        int power = 1;
        result += i + 1;
        for (int j = 0; filecontent[i][j] != '\0'; j++) {
            if (find_substring("red\0", filecontent[i], j) == 1) {
                int temp = j - 2;
                int num = 0;
                while (temp >= 0) {
                    if (filecontent[i][temp] == ' ')
                        break;
                    num += (filecontent[i][temp] - '0') * power;
                    power *= 10;
                    temp--;
                }
                if (num > max_red) {
                    result -= i + 1;
                    break;
                }
                power = 1;
            }
            if (find_substring("green\0", filecontent[i], j) == 1) {
                int temp = j - 2;
                int num = 0;
                while (temp >= 0) {
                    if (filecontent[i][temp] == ' ')
                        break;
                    num += (filecontent[i][temp] - '0') * power;
                    power *= 10;
                    temp--;
                }
                if (num > max_green) {
                    result -= i + 1;
                    break;
                }
                power = 1;
            }
            if (find_substring("blue\0", filecontent[i], j) == 1) {
                int temp = j - 2;
                int num = 0;
                while (temp >= 0) {
                    if (filecontent[i][temp] == ' ')
                        break;
                    num += (filecontent[i][temp] - '0') * power;
                    power *= 10;
                    temp--;
                }
                if (num > max_blue) {
                    result -= i + 1;
                    break;
                }
                power = 1;
            }
        }
    }
    printf("%d\n", result);
    for (int i = 0; filecontent[i] != NULL; i++)
        free(filecontent[i]);
    free(filecontent);
    return 0;
}