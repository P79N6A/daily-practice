defmodule Main do
  def anagram?(word1, word2)
  when byte_size(word1) == byte_size(word2) do
    eq?(word1, word2, 0)
  end

  def anagram?(_, _) do
    false
  end

  defp eq?(word1, word2, i) do
    n = byte_size(word1)
    if i == n do
      true
    else
      if String.at(word1, i) != String.at(word2, n - i - 1)  do
        false
      else
        eq?(word1, word2, i+1)
      end
    end
  end
end

# true
Main.anagram?("heleh", "heleh")
  |> IO.puts
# false
Main.anagram?(123, 'ehlo')
  |> IO.puts
