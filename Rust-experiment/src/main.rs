use rand::Rng;
use std::cmp::Ordering;
use std::time::Instant;

struct A<T: Ord, U: Ord> {
    first: T,
    second: U,
    third: i64,
}

impl<T: Ord, U: Ord> Ord for A<T, U> {
    fn cmp(&self, other: &Self) -> Ordering {
        match self.first.cmp(&other.first) {
            Ordering::Equal => match self.second.cmp(&other.second) {
                Ordering::Equal => self.third.cmp(&other.third),
                other => other,
            },
            other => other,
        }
    }
}

impl<T: Ord, U: Ord> Eq for A<T, U> {}

impl<T: Ord, U: Ord> PartialEq for A<T, U> {
    fn eq(&self, other: &Self) -> bool {
        self.first == other.first && self.second == other.second && self.third == other.third
    }
}

impl<T: Ord, U: Ord> PartialOrd for A<T, U> {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

#[derive(PartialOrd, PartialEq)]
struct B {
    first: String,
    second: i32,
    third: i64,
}

impl Ord for B {
    fn cmp(&self, other: &Self) -> Ordering {
        if self.first.eq(&other.first) {
            if self.second == other.second {
                return self.third.cmp(&other.third);
            }
            return self.second.cmp(&other.second);
        }
        self.first.cmp(&other.first)
    }
}

impl Eq for B {}

fn test(mut arr: Vec<B>) -> u128 {
    let start = Instant::now();
    arr.sort();
    let elapsed = start.elapsed();
    let elapsed_micros = elapsed.as_micros();
    elapsed_micros
}

fn generics_test<T: Ord, U: Ord>(mut arr: Vec<A<T, U>>) -> u128 {
    let start = Instant::now();
    arr.sort();
    let elapsed = start.elapsed();
    let elapsed_micros = elapsed.as_micros();
    elapsed_micros
}

fn generate_random_b() -> B {
    let mut rng = rand::thread_rng();
    let random_number: u128 = rng.gen();
    let first: String = format!("{}", random_number);
    let second: i32 = rng.gen();
    let third: i64 = rng.gen();
    B {
        first,
        second,
        third,
    }
}

fn main() {
    let mut total_sort_time_for_a: u128 = 0;
    let mut total_sort_time_for_b: u128 = 0;

    let n_experiments = 10;

    for _ in 0..n_experiments {
        let mut vec_b: Vec<B> = Vec::with_capacity(100000);
        for _ in 0..100000 {
            let random_b = generate_random_b();
            vec_b.push(random_b);
        }

        let mut vec_a: Vec<A<String, i32>> = Vec::with_capacity(100000);
        for _ in 0..100000 {
            let random_b = generate_random_b();
            let random_a = A {
                first: random_b.first,
                second: random_b.second,
                third: random_b.third,
            };
            vec_a.push(random_a);
        }

        let sort_time_for_a = generics_test(vec_a);
        let sort_time_for_b = test(vec_b);

        total_sort_time_for_a += sort_time_for_a;
        total_sort_time_for_b += sort_time_for_b;
    }

    println!(
        "Mean elapsed time of sort with generics: {} mcs",
        total_sort_time_for_a / n_experiments
    );
    println!(
        "Mean elapsed time of sort without generics: {} mcs",
        total_sort_time_for_b / n_experiments
    );
}
