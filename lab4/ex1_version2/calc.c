#include <stdbool.h>

double calc(double width, double height, int chambers, int material, bool sill) {
    double area = width * height;
    double price = 0;

    double table[3][3] = {
        {0, 2.5, 3},
        {0, 0.5, 1},
        {0, 1.5, 2}
    };

    price = area * table[material][chambers];

    if (sill) price += 350;

    return price;
}
