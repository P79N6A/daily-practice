use std::marker::PhantomData;
use std::mem::align_of;

struct RefWithFlag<'a, T: 'a> {
    ptr_and_bit: usize,
    behaves_like: PhantomData<&'a T>
}

impl<'a, T: 'a> RefWithFlag<'a, T> {
    fn new(ptr: &'a T, flag: bool) -> RefWithFlag<T> {
        assert!(align_of::<T>() % 2 == 0);
        RefWithFlag {
            ptr_and_bit: ptr as *const T as usize | flag as usize,
            behaves_like: PhantomData
        }
    }

    fn get_ref(&self) -> &'a T {
        unsafe {
            let ptr = (self.ptr_and_bit & !1) as *const T;
            &*ptr
        }
    }

    fn get_flag(&self) -> bool {
        self.ptr_and_bit & 1 != 0
    }
}

pub fn demo_ref_with_flag() {
    let vec = vec![10, 20, 30];
    let flagged = RefWithFlag::new(&vec, true);
    println!("[RefWithFlag][1]: {}, flag: {}", flagged.get_ref()[1], flagged.get_flag());
}