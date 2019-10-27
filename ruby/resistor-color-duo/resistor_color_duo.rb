#!/usr/bin/env ruby -w

class ResistorColorDuo
	RESISTOR_COLORS = %w[
		black brown red orange yellow green blue violet grey white
		].freeze
	def self.value(args=nil)
		args
		.take(2)
		.map {|color| RESISTOR_COLORS
		.find_index(color)}
		.join
		.to_i
	end
end
