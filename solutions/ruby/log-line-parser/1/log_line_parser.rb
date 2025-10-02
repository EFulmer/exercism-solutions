class LogLineParser
  def initialize(line)
    @line = line
    @message = @line.split(":")[1].strip
    @log_level = @line.split(":")[0][1..-2].downcase
  end

  def message
    @message
  end

  def log_level
    @log_level
  end

  def reformat
    "#{@message} (#{@log_level})"
  end
end
