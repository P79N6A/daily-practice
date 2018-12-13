defmodule Main do
  def main do
    filename = "/etc/hosts"
    case File.open(filename) do
      { :ok, file } ->
        IO.puts "the whole #{filename}:\n#{IO.read(file, :all)}"
      { :error, reason } ->
        IO.puts "Failed to open file: #{{reason}}"
    end
  end
end

Main.main
