defmodule Main do
  def list do
    ["now", "is", "the", "time"]
  end

  def main do
    longest = Enum.reduce(Main.list, fn word, longest ->
      if String.length(word) > String.length(longest) do
        word
      else
        longest
      end
    end)
    IO.puts "longest: #{longest}"
  end
end

Main.main
