defmodule Main do
  def is_prime_i(i, upper, x) do
    if i > upper do
      true
    else
      if Integer.mod(x, i) == 0 do
        false
      else
        is_prime_i(i+1, upper, x )
      end
    end
  end

  def is_prime(x) do
    i = 2
    upper = round(:math.sqrt(x))
    # IO.puts "x: #{x}, upper: #{upper}"
    is_prime_i(i, upper, x)
  end

  def span(n) do
    for x <- Enum.into(2..n, []), is_prime(x), do: x
  end

end

n = 100
IO.puts "2~#{n} 以内的素数："
IO.inspect Main.span(n)
