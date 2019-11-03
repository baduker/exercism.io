#!/usr/bin/env ruby -w

class HighScores

  attr_reader :scores

  def initialize(scores)
    @scores = scores
  end

  def latest
    scores[-1]
  end

  def personal_best
    scores.max
  end

  def personal_top_three
    scores.sort.reverse[0...3]
  end

  def latest_is_personal_best?
    scores[-1].eql? scores.max
  end

end
