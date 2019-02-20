# 4. Monoidal Category

## Definition
A monoidal category consists of a category C equipped with a functor:
```
    ⊗ : C × C −→ C
```
and a distinguished object I, together with natural isomorphisms,
```
    αABC : A ⊗ (B ⊗ C) ⥲ (A ⊗ B) × C,
    
    λA : I ⊗ A ⥲ A,      ρA : A ⊗ I ⥲ A.
```
