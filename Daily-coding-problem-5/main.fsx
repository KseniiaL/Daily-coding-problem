//A pair type
type Val = Pair of (int * int)

//Given function, that constructs a pair
let cons a b: Val =
    let pair() =
        Pair(a,b)

    pair()

//Simple implementation of car
let car f: int =
    match f with
    | Pair(a,b) -> a

//Simple implementation of cdr
let cdr f: int =
    match f with
    | Pair(a,b) -> b

car(cons(1,3))
cdr(cons(3,4))
car(cons(cdr(cons(4,5)), car(cons(2,5))))