class Acronym

  attr_reader :sentence

  def self.abbreviate(sentence="")
    sentence.split(/[\s\-]/).map {|word| word[0]}.join.upcase
  end

end
