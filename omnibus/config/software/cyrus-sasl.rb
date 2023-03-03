name "cyrus-sasl"
default_version "2.1.28"

dependency "gdbm"

if redhat?
    dependency "e2fsprogs"
end

source :url => "https://github.com/cyrusimap/cyrus-sasl/releases/download/cyrus-sasl-#{version}/cyrus-sasl-#{version}.tar.gz",
       :sha256 => "7ccfc6abd01ed67c1a0924b353e526f1b766b21f42d4562ee635a8ebfc5bb38c",
       :extract => :seven_zip

relative_path "cyrus-sasl-#{version}"

build do

  env = {
    "LDFLAGS" => "-Wl,-rpath,#{install_dir}/embedded/lib -L#{install_dir}/embedded/lib",
    "CFLAGS" => "-L#{install_dir}/embedded/lib -I#{install_dir}/embedded/include",
    "LD_RUN_PATH" => "#{install_dir}/embedded/lib",
  }

  configure_command = ["./configure",
                        "--prefix=#{install_dir}/embedded",
                        "--with-gdbm=#{install_dir}/embedded",
                        "--with-dblib=gdbm"]

  command configure_command.join(" "), env: env
  command "make", :env => env
  command "make install", :env => env

end
