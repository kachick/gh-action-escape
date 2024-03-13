size = ARGV[0].slice(/\n*\z/).size
raise size.to_s if size != Integer(ARGV[1])
