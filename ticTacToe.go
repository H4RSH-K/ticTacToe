package main

import ("fmt")

func main(){
  gameOver := false
  board := [9]int{0,0,0,0,0,0,0,0,0}
  turNum := 1

  for gameOver != true{
    Board(board)
    player := 0
    if turNum % 2 == 1{
      fmt.Println("Player 1's turn")
      player = 1
    } else {
      fmt.Println("Player 2's turn")
      player = 2
    }

    Move := moveQuery()
      if Move == 9 {
        return
      }
    board = executeMove(Move, player, board)

    result := whoWon(board)
    if result > 0{
      fmt.Printf("Player %d wins!\n\n", result)
      gameOver = true
    } else if turNum == 9 {
      fmt.Printf("Tie game!\n\n")
      gameOver = true
    } else {
      turNum++
    }
  }

}

func moveQuery() int{
  fmt.Println("Select a move")
  var moveInt int
  fmt.Scan(&moveInt)
  return moveInt
}

func executeMove(moveInt int, player int, b [9]int) [9]int {

  if b[moveInt] != 0 {
    fmt.Println("Please pick an unoccupied space.")
    moveInt = moveQuery()
    b = executeMove(moveInt, player, b)
  } else {
    if player == 1 {
      b[moveInt] = 1
    } else if player == 2 {
      b[moveInt] = 10
    }
  }

  for moveInt > 9 {
      fmt.Println("Please enter a number under 10.")
      moveInt = moveQuery()
  }

  if player == 1{
    b[moveInt] = 1
  }else if player == 2{
    b[moveInt] = 10
  }
  return b
}

func Board(b [9]int) {
  for i, v := range b {
    if v == 0{
      fmt.Printf("%d", i)
    } else if v == 1{
      fmt.Printf("X")
    } else if v == 10{
      fmt.Printf("O")
    }
    if i> 0 && (i+1) % 3 == 0{
      fmt.Printf("\n")
    } else {
      fmt.Printf(" | ")
    }
  }
}

func whoWon(b [9]int) int {
  sums := [8] int {0,0,0,0,0,0,0,0}

  sums[0] = b[2]+b[4]+b[6]
  sums[1] = b[0]+b[3]+b[6]
  sums[2] = b[1]+b[4]+b[7]
  sums[3] = b[2]+b[5]+b[8]
  sums[4] = b[0]+b[4]+b[8]
  sums[5] = b[6]+b[7]+b[8]
  sums[6] = b[3]+b[4]+b[5]
  sums[7] = b[0]+b[1]+b[2]
  for _, v := range sums {
    if v == 3{
      return 1
    } else if v == 30{
      return 2
    }
  }
  return 0
}
