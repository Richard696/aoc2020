


fn main(){
    // read file and throw err if encountered
    let input = std::fs::read_to_string("input.txt").unwrap();

    // split the input by new line, only take in not empty value, parse as int and map into array
    let mut xs = input.split("\n").filter(|s| !s.is_empty()).map(|x| x.parse::<i32>().unwrap()).collect::<Vec<_>>();

    let n = xs.len();
    // Time complexity: O(N^2)
    // part 1
    // for i in 0..n - 1{
    //     for j in i + 1..n{
    //         if xs[i] + xs[j] == 2020{
    //             println!("{}", xs[i]*xs[j]);
    //             return;
    //         }
    //     }
    // }

    // part 2
    // O(N^3)
    // for i in 0..n - 2{
    //     for j in i+1..n - 1 {
    //         for k in j+1..n {
    //             if xs[i] + xs[j] + xs[k]== 2020{
    //                 println!("{}", xs[i]*xs[j]*xs[k]);
    //                 break;
    //         }
    //         }
    //     }
    // }

    // Time complexity O(N Log N)
    // Part 1
    xs.sort(); // sort the array
    // for i in 0..n{ //O(N)
    //     if let Ok(j) = xs.binary_search(&(2020-xs[i])) {
    //         if i != j{
    //             println!("{}", xs[i]*xs[j]);
    //             break;
    //         }
    //     }
    // }
    
    // Part 2
    // O(N^2Log(N))
    for i in 0..n-1{ //O(N)
        for j in i+1..n{
            if let Ok(k) = xs.binary_search(&(2020-xs[i]-xs[j])) {
                if j != k{
                    println!("{}", xs[i]*xs[j]*xs[k]);
                    break;
                }
            }
        }
    }
        
}