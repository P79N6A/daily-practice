defmodule Fns do
  def add(a, b) do
    a + b
  end

  def sub(a, b) do
    a - b
  end

  def mul(a, b) do
    a * b
  end

  def div(a, b) do
    a / b
  end
end

IO.puts Fns.add(4, 5)
          |> Fns.sub(3)
          |> Fns.mul(4)
          |> Fns.div(2)
          |> Fns.add(20)
