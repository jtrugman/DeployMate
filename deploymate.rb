class Deploymate < Formula
  desc "An opinionated simple deployment system"
  homepage "https://github.com/jtrugman/deploymate"
  url "https://github.com/jtrugman/deploymate/archive/refs/tags/v0.1.0.tar.gz"
  sha256 "ba7de75e37f3180d10ed3faf0925ec89dd82e5fb613c00c754d4859d1ebb9ee9"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w")
  end

  test do
    assert_match "DeployMate", shell_output("#{bin}/deploymate --help")
  end
end
