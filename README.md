# expense-tracker

https://roadmap.sh/projects/expense-tracker

---

## Run

```sh
go run main.go [commands]
```

## Build

```sh
go build
```

## Example Commands

### Adding a new expense

```sh
./expense-tracker add --description "Lunch" --amount 20
```
```sh
./expense-tracker add --description "Dinner" --amount 10
```

### Listing all expenses

```sh
./expense-tracker list
```

### Expense summary

```sh
./expense-tracker summary
```
```sh
./expense-tracker summary --month 8
```

### Deleting an expense

```sh
./expense-tracker delete --id 2
```
