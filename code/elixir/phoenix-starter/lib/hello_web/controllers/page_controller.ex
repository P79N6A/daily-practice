defmodule HelloWeb.PageController do
  use HelloWeb, :controller

  def index(conn, _params) do
    render conn, "index.html"
  end

  def greeting(conn, _params) do
    render conn, "greeting.html"
  end

  def hello(conn, _params) do
    text conn, "hello, buddy!"
  end

  def holy(conn, _params) do
    json conn, %{greeting: "holy crap!"}
  end
end
