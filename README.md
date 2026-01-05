# 3.4.1.Goroutines_and_channels
# Paralēlā masīva kārtošana ar gorutīnām 

Programma nolasa masīva elementu skaitu no tastatūras; ja tas ir mazāks par 10, tiek izvadīts kļūdas paziņojums un programma beidz darbu.<br>
<br>
Pēc tam lietotājs ievada masīva elementus vai izvēlas aizpildīt masīvu ar nejaušiem veseliem skaitļiem.<br>
<br>
Masīvs tiek sadalīts N aptuveni vienādās daļās, kur N tiek aprēķināts kā kvadrātsakne no elementu skaita, bet N nedrīkst būt mazāks par 4.<br>
<br>
Katra masīva daļa tiek sakārtota augošā secībā atsevišķā gorutīnā (neizmantojot Go standarta bibliotēkas gatavās kārtošanas funkcijas), un katra gorutīna izvada savu apakšmasīvu pirms un pēc sakārtošanas.<br>
<br>
Pēc visu gorutīnu pabeigšanas galvenā gorutīna apvieno sakārtotās daļas vienā pilnībā sakārtotā masīvā un izvada gala rezultātu.<br>
