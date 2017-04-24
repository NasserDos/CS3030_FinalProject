#!/usr/bin/perl

use strict;
use warnings;
use File::Path;
use LWP::Simple;


sub copyFile{
(my $sec,my $min,my $hour,my $mday,my $mon,my $year,my $wday,my $yday,my $isdst) = localtime();

my $cmd='scp';
my $from='nb19445@icarus.cs.weber.edu:/home/hvalle/submit/cs3030/files/FRED.csv';
my $to = sprintf("fredData/%02d",$mon);

system($cmd,$from,$to);

}


sub prepStruct{

    print "Preparing file structure";
    foreach my $i ("01".."12") {
        mkpath "fredData/$i" 
    }
}



sub main{

    prepStruct();
    copyFile();
}





main()


