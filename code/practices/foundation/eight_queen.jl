function find_safe_coord(board::Array{Int}, curr_row, n; trace_back=false)
    start = 1
    if trace_back
        start = board[curr_row] + 1
    end
    for col = start : n
        has_menace = false
        for row = 1 : curr_row - 1
            prev = board[row]
            if prev == col || curr_row - row  == abs(col - prev)
                has_menace = true
            end
        end
        !has_menace && return col
    end
    nothing
end

function n_queens_puzzle(n)
    board = zeros(Int, n)
    curr_row = 1
    while curr_row <= n
        coord = find_safe_coord(board, curr_row, n)
        if coord != nothing
            board[curr_row] = coord
            curr_row == n && produce(copy(board))
        end
        if coord == nothing || curr_row == n
            while true
                curr_row -= 1
                curr_row == 0 && return produce(nothing)
                coord = find_safe_coord(board, curr_row, n, trace_back=true)
                if coord != nothing
                    board[curr_row] = coord
                    break
                end
            end
        end
        curr_row += 1
    end
end

# 逐行输出[x,y]
solutions = @task n_queens_puzzle(8)
for solution in solutions
    if solution != nothing
        for row in eachindex(solution)
            print("[$row,$(solution[row])] ")
        end
        println()
    end
end
