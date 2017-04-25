#!/usr/bin/perl

use strict;
use warnings;
use File::Path;
use Getopt::Std;
#use LWP::Simple;


sub copyFile{
=pod
    subroutine that copied the file from icarus.
    it creates a date array of the current date, takes the required
    date information to stamp it onto the file

    this subroutine also takes the parameters from the Getopts
=cut

(my $sec,my $min,my $hour,my $mday,my $mon,my $year,my $wday,my $yday,my $isdst) = localtime();

my $cust = shift @_;
my $dataFile = shift @_;

my $cmd='scp';
my $timestamp = sprintf("%04d-%02d-%02d",$year+1900,$mon+1,$mday);
my $from=sprintf('nb19445@icarus.cs.weber.edu:/home/hvalle/submit/cs3030/files/%s',$dataFile);
my $to = sprintf("%s/%02d/FRED.csv.%04d-%02d-%02d",$cust,$mon+1,$year+1900,
    $mon+1,$mday);

print "Checking icarus for the file\n";
 my $rc = system "ssh",sprintf('nb19445@icarus.cs.weber.edu'),
 "test","-e","/home/hvalle/submit/cs3030/files/$dataFile";

 if($rc){
print "File $dataFile was not found\n";
print "Failed to scp the file\n";
 }else{
print "$dataFile was found\n";
print "Checking Structure...\n";
prepStruct();
print "Getting file from icarus\n";
system($cmd,$from,$to);
print "scp was successful\n";

printf("File located at [fredData/%02d/FRED.csv.%s]\n",$mon+1,$timestamp);
 }


}


sub prepStruct{
=pod
    sub routine that handles the file structure creation
    it creates the structure if it didn't exist
=cut
    print "Preparing file structure\n";
    foreach my $i ("01".."12") {
        mkpath "fredData/$i";
    }
}


sub checkusage{
=pod
    Checks the user input for correct arguments
    Returns 0 if the user fails to have the correct options
    Returns 1 if the user passes the correct options
=cut
    my $opts = shift;

    my $c = $opts->{"c"};
    my $f = $opts->{"f"};

    unless(defined($c) and defined($f)){
        return 0;
    }

    return 1;


}

sub Usage{
=pod
    The help subroutine, shows how the script is used

=cut
    my $help = <<USAGE;
    Usage: nasser_binshabeeb.hw8 [-c customerDataFolder] [-f dataFile]
    Both arguments are required
USAGE
    die $help;
}


sub main{
=pod
    The main subroutine
    gets the options and runs the other subroutines
    calls the check options sub as well.
=cut

    my %opts;
    getopts('c:f:', \%opts) ;

    if(!checkusage(\%opts)){
    Usage();
    }

    copyFile($opts{'c'},$opts{'f'});
    exit(0);

}


# Main function call
main()

