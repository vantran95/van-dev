# Factory Pattern

**Factory Pattern is a creational design pattern that provides a method for creating objects without specifying their exact concrete class. It helps reduce dependency on concrete classes, increases flexibility, and makes it easier to extend when there are requirements to change or add new different classes.**

**In Go, because there are no classes as in other object-oriented languages, we can use interfaces combined with factory functions to implement the Factory Pattern.**

- **Implementation in Go:**

    ```go
    package main
    
    import "fmt"
    
    type Product interface {
        Use() string
    }
    
    type ConcreteProductA struct{}
    func (p *ConcreteProductA) Use() string { return "ProductA" }
    
    type ConcreteProductB struct{}
    func (p *ConcreteProductB) Use() string { return "ProductB" }
    
    type Creator interface {
        CreateProduct() Product
    }
    
    type ConcreteCreatorA struct{}
    func (c *ConcreteCreatorA) CreateProduct() Product { return &ConcreteProductA{} }
    
    type ConcreteCreatorB struct{}
    func (c *ConcreteCreatorB) CreateProduct() Product { return &ConcreteProductB{} }
    
    func main() {
        creators := []Creator{&ConcreteCreatorA{}, &ConcreteCreatorB{}}
        for _, creator := range creators {
            product := creator.CreateProduct()
            fmt.Println(product.Use())
        }
    }
    
    ```

---

### Factory Pattern in Go: Case Study

The **Factory Pattern** is a **creational design pattern** that provides a method for initializing objects without needing to specify their exact struct. It helps reduce dependency on specific classes, increases flexibility, and makes it easier to extend the codebase when adding or changing different classes.

In Go, since there are no classes (unlike object-oriented languages), you can use **interfaces** combined with **factory functions** to implement the Factory Pattern.


### Explanation:

- **Vehicle Interface**: This acts as the common interface for all vehicle types. Specific vehicles like `Car`, `Bicycle`, and `Truck` each implement the `Drive()` method in their own way.
- **Factory Method**: The `VehicleFactory()` function acts as a factory. Based on the value of `vehicleType`, it initializes and returns the corresponding object (`Car`, `Bicycle`, or `Truck`).
- **Flexibility**: When you need to add a new vehicle type (e.g., `Motorcycle`), you only need to add a new `case` in `VehicleFactory()` without changing other parts of the code.
- **Error Handling**: If the user requests a non-existent vehicle type, the factory returns `nil` and you can handle the error accordingly.

---

### When to Use the Factory Pattern?

- **Complex Object Creation**: When the process of creating objects is complex, you can hide the initialization logic inside the factory, making the code easier to read and maintain.
- **Reducing Dependency on Specific Classes**: The Factory Pattern helps decouple your code from specific classes, relying on interfaces instead, making it easier to extend and maintain.
- **High Scalability**: When you need to extend the system by adding new types of objects, you only need to modify the factory without changing the existing code.

---
### When NOT to Use the Factory Pattern?

- **Simple Requirements**: If your object is very simple and does not require any complex initialization logic, using the Factory Pattern may add unnecessary complexity to your code.
- **Processing Overhead**: In some cases, initializing objects through a factory can introduce overhead, especially when objects need to be created frequently without complex logic.

---

### Conclusion:

The Factory Pattern is an efficient way to create objects without specifying their concrete classes. In Go, you can implement the Factory Pattern using interfaces and factory functions.

The **vehicle management system case study** demonstrates how the Factory Pattern helps manage different vehicle types without being dependent on the specific initialization logic of each type, making the code easier to maintain and extend.

The Factory Pattern is especially useful in large systems where object creation is complex, and there are multiple types of objects to manage.