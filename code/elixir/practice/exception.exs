defmodule Main do

  defp _handle(reason) do
    IO.puts "error occurs: #{reason}"
  end

  defp _readfile(file) do
    case IO.read(file, :line) do
      :eof ->
        IO.puts "\nEND\n"
      {:error, reason} ->
        _handle(reason)
      data ->
        IO.write data
        _readfile(file)
    end
  end

  def open(filename) do
    case File.open(filename) do
      {:ok, file} ->
        _readfile(file)
      {:error, reason} ->
        _handle(reason)
    end
  end

end

Main.open("/etc/hosts")
