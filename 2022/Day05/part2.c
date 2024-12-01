#include <fcntl.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>
#include <sys/stat.h>
#include <stddef.h>
#include <string.h>

typedef struct crates_s {
    char crate_name;
    struct crates_s *next;
} crates_t;

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

crates_t *add_end_node(crates_t *stacks, char crate, int stack)
{
    crates_t *result = malloc(sizeof(crates_t));
    crates_t *cursor = stacks;
    result->crate_name = crate;
    result->next = NULL;
    if (stacks == NULL) {
        return result;
    }
    while (cursor->next != NULL)
        cursor = cursor->next;
    cursor->next = result;
    return stacks;
}

crates_t **init_map(char **filecontent)
{
    crates_t **result = NULL;
    int num_stack = 0;
    for (int i = 0; filecontent[7][i] != '\0'; i++)
        if (filecontent[7][i] == '[')
            num_stack += 1;
    result = malloc(sizeof(crates_t *) * (num_stack + 1));
    result[num_stack] = NULL;
    int k = 1;
    for (int i = 0; i < num_stack; i++) {
        result[i] = NULL;
        for (int j = 0; strcmp(filecontent[j + 1], "\0") != 0; j++) {
            if (filecontent[j][k] != ' ')
                result[i] = add_end_node(result[i], filecontent[j][k], i + 1);
        }
        k += 4;
    }
    return result;
}

void print_stacks(crates_t **orga, int len)
{
    for (int i = 0; i < len; i++) {
        crates_t *cursor = orga[i];
        printf("%d : ", i + 1);
        while (cursor != NULL) {
            printf("[%c]", cursor->crate_name);
            cursor = cursor->next;
        }
        printf("\n");
    }
    printf("\n");
}

void free_stack(crates_t *stack)
{
    if (stack->next == NULL) {
        free(stack);
        return;
    }
    free_stack(stack->next);
    free(stack);
    return;
}

int get_value(char const *line, int index)
{
    int result = 0;
    int power = 0;
    int temp = 1;
    for (int i = index; line[i] != '\0'; i++) {
        if (line[i] < '0' || line[i] > '9')
            break;
        power += 1;
    }
    for (int i = index; line[i] != '\0'; i++) {
        if (line[i] < '0' || line[i] > '9')
            break;
        for (int j = 0; j < power - 1; j++) {
            temp *= 10;
        }
        result += (line[i] - '0') * temp;
        temp = 0;
    }
    return result;
}

int main(void)
{
    char **filecontent = load_2d_arr_from_file(load_file_in_mem());
    crates_t **orga = init_map(filecontent);
    int num_stack = 0;
    for (int i = 0; filecontent[7][i] != '\0'; i++)
        if (filecontent[7][i] == '[')
            num_stack += 1;
        print_stacks(orga, num_stack);

    for (int i = 10; filecontent[i] != NULL; i++) {
        int value[3] = {0, 0, 0};
        int index = 0;
        for (int j = 0; filecontent[i][j] != '\0'; j++) {
            if (filecontent[i][j] >= '0' && filecontent[i][j] <= '9') {
                value[index] += get_value(filecontent[i], j);
                if (filecontent[i][j + 1] < '0' || filecontent[i][j + 1] > '9')
                    index++;
            }
        }
        crates_t *temp_start = orga[value[1] - 1];
        crates_t *temp_start2 = orga[value[1] - 1];
        for (int j = 0; j < value[0] - 1; j++) {
            temp_start = temp_start->next;
            temp_start2 = temp_start2->next;
        }
        temp_start = temp_start->next;
        temp_start2->next = orga[value[2] - 1];
        orga[value[2] - 1] = orga[value[1] - 1];
        orga[value[1] - 1] = temp_start;
        print_stacks(orga, num_stack);
    }
    for (int i = 0; orga[i] != NULL; i++)
        printf("%c", orga[i]->crate_name);
    for (int i = 0; orga[i] != NULL; i++)
        free_stack(orga[i]);
    free(orga);
    for (int i = 0; filecontent[i] != NULL; i++)
        free(filecontent[i]);
    free(filecontent);
    return 0;
}