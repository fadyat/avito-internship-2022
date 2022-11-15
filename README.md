## Service for working with user balances

> Created for the [Avito Internship](https://github.com/avito-tech/internship_backend_2022)

### Todo:

- Метод начисления средств на баланс.

> Принимает `user_id` и `amount`.

- Метод резервирования средств с основного баланса на отдельном счете.

> Принимает `user_id`, `service_id`, `order_id`, `amount`.

- Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии.

> Принимает `user_id`, `service_id`, `order_id`, `amount`.

- Метод отмены резервирования – отменяет резервирование средств.

> Принимает `user_id`, `service_id`, `order_id`, `amount`.


- Метод получения баланса пользователя.

> Принимает `user_id`.

- Валидация входных данных.

- Покрытие тестами.

- Документация.

- Бухгалтерия раз в месяц просит предоставить сводный отчет по всем пользователем, с указанием сумм выручки по каждой из
  предоставленной услуги для расчета и уплаты налогов.

> Реализовать метод для получения месячного отчета. На вход: год-месяц. На выходе ссылка на CSV файл.

- Получение списка транзакций по пользователю.

> Реализовать метод для получения списка транзакций по пользователю. На вход: `user_id`, `limit`, `offset`.
> На выходе: список транзакций.
>
> Добавить сортировку по дате и сумме.