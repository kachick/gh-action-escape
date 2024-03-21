size = ARGV[0].slice(/\n*\z/m).size
raise size.to_s if size != Integer(ARGV[1])
