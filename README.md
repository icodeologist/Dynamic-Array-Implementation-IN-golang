- What is dynamic array
  - Array which can add any amount of elements
  - When the limitation of the size is reached it gives itself double free size and also copies all the old entries
  - It also frees the old memory

Python list()
- append() -> Add elements at the end of the last element  usually at n  where n = size of the current array
- a[[i]] -->  Access the element at ith index where 0>= i < n, n = len(array)
- a[[i]] = n -->  Change the element at the given index of array. But make it same type
- delete a middle key -->  Deleting any item in the middle index should move elements accordingly. So if I delete element at index 2 then all the left elements should move right and free the last place


Some Slice Things:
- Var declares zero value to a variable
- ie... var s []string
-  if s == nil: is True
- Becuase s is nil value equalled to a slice
- Nil Slice -> Slice pointers in not pointing to a underlying array
- Empty Slice -> Len(slice) elements are empty
- Make(type, length , cap)
