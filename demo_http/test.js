var rand = function() {
    return Math.random().toString(36).substr(2);
};

export const random = function() {
    return rand() + rand();
};