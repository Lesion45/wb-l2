# Паттерн [«Цепочка вызовов»](https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern)
### Суть паттерна
Цепочка вызовов — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по 
цепочке обработчиков. Каждый обработчик решает, может ли он обработать запрос, или передаёт его дальше по цепочке.
### Когда применяем?
1. В разрабатываемой системе имеется группа объектов, которые могут обрабатывать сообщения определенного типа.
2. Все сообщения должны быть обработаны хотя бы одним объектом системы.
### Плюсы
* Легко добавлять новые обработчики или менять порядок цепочки.
* Разделение логики обработки на независимые сущности.
### Минусы
* Сложность отладки.