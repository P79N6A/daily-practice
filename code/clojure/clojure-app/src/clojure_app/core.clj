(ns clojure-app.core
  (:gen-class)
  (:require [clojure.core.reducers :as r]))

(def numbers (into [] (range 0 10000000)))

(def counts {"apple" 2 "orange" 1})

(def pages ["one potato two potato three potato four"
            "five potato six potato seven potato more"])

(defn get-words [text] (re-seq #"\w+" text))

(defn parallel-sum [numbers]
  (r/fold + numbers))

(defn count-words_parallel [pages]
  (reduce (partial merge-with +)
          (pmap #(frequencies (get-words %)) pages)))

(defn -main
  "I don't do a whole lot ... yet."
  [& args]
  (println "Hello, World!"))
