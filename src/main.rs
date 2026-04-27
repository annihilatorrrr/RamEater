use std::env;
use std::process::{exit, Command};
use std::thread::sleep;
use std::time::Duration;
use sysinfo::System;

fn ram_to_consume() -> Option<Vec<u8>> {
    let mut system = System::new_all();
    system.refresh_memory();
    let total_memory = system.total_memory();
    let used_memory = system.used_memory();
    let take: f64 = env::var("TAKE")
        .ok()
        .and_then(|v| v.parse::<u32>().ok())
        .unwrap_or(15) as f64
        / 100.0;
    let memory_to_allocate = (total_memory as f64 * take) - used_memory as f64;
    if memory_to_allocate > 0.0 {
        let allocated_mem = vec![0_u8; memory_to_allocate as usize];
        println!(
            "Allocated {} bytes to reach target memory usage.",
            allocated_mem.len()
        );
        Some(allocated_mem)
    } else {
        println!("No need to allocate memory.");
        None
    }
}

fn main() {
    let _allocated_mem = ram_to_consume();
    println!("Done!");
    loop {
        sleep(Duration::from_secs(24 * 3600));
        if env::var("NOCPUB").is_err() {
            let mut result: u64 = 1;
            for i in 1u64..1_000_000 {
                result = result.wrapping_mul(i);
            }
        }
        let current_exe = env::current_exe().expect("Failed to get current executable path");
        Command::new(current_exe)
            .status()
            .expect("Failed to re-execute process");
        exit(0);
    }
}
