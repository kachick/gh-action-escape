size = ARGV[1].slice(/\n*\z/).size
raise size.to_s if size != Integer(ARGV[2])
