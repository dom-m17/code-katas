// Timmy & Sarah think they are in love, but around where they live, they will only know once they pick a flower each. 
// If one of the flowers has an even number of petals and the other has an odd number of petals it means they are in love.

// Write a function that will take the number of petals of each flower and return true if they are in love and false if they aren't.

function lovefunc(flower1: number, flower2: number): boolean {
    return Math.abs(flower1 - flower2) % 2 === 1
}

console.log(lovefunc(1, 4))
console.log(lovefunc(2, 2))
console.log(lovefunc(0, 1))
console.log(lovefunc(0, 0))

// https://www.codewars.com/kata/555086d53eac039a2a000083/typescript

export function lovefunc2(flower1: number, flower2: number): boolean {
    return flower1 % 2 != flower2 % 2;
}

export function lovefunc3(flower1: number, flower2: number): boolean {
    return (flower1 + flower2) % 2 === 1;
}

// All solutions were roughly the same but these two may be slightly better because they don't require another package