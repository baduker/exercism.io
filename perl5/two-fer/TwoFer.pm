package TwoFer;
use strict;
use warnings;
use Exporter 'import';
our @EXPORT_OK = qw(two_fer);
use experimental qw(signatures);

sub two_fer ($name = undef) {
    my ($name) = @_;
    $name = "you" unless $name;
    "One for $name, one for me.";
}

1;
