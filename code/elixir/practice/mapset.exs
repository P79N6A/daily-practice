defmodule Main do
  def new_set(range) do
    Enum.into range, MapSet.new
  end

  def main do
    set1 = Main.new_set(1..5)
    set2 = Main.new_set(3..8)
    IO.inspect MapSet.difference set1, set2
    IO.inspect MapSet.union set1, set2
    IO.inspect MapSet.intersection set1, set2
  end
end

Main.main
