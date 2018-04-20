require 'bundler'
Bundler.setup()
require 'pry'

namespace :xdr do

  task :update => [:generate]

  task :generate do
    require "pathname"
    require "xdrgen"
    require 'fileutils'
    FileUtils.rm_f("xdr/xdr_generated.go")

    paths = Pathname.glob("xdr/raw/*.x").sort
    compilation = Xdrgen::Compilation.new(
      paths,
      output_dir: "xdr",
      namespace:  "xdr",
      language:   :go
    )
    compilation.compile
    system("gofmt -w xdr/xdr_generated.go")
  end
end
