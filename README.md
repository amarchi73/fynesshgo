# fynesshgo

### Una Gui per gestire le connessoni SSH con Go e Fyne
**_usa pass per eventualmente gestire le password_**

Per lavoro (e svago) mi capita di dovermi collegare in ssh a un discreto numeo di macchine, tenere a mente
i vari utenti e IP e fare copia incolla tutte le volte era una rottura.

Per le password usavo già il comando pass quindi quelle sono a posto e ho pensato che sarebbe stato interessante 
integrarlo con un sistema che mi consentisse di richiamare la connessione ssh con un semplice parametro e inserire 
poi la password GPG.

Con lo stesso criterio integrerò una login per l'utilizzo del programma e la cifratura PGP dei JSON 
che contengono i dati di connessione. 

Questo è il mio primo progetto "completo" usando Go e Fyne, quindi probabilmente parecchie cose sono scritte con il c**o, 
quando avrò implementato le funzionalità che mi interessano vedrò di fare un po' di refactoring come ad esempio dare dei nomi
decenti alle variabili e usare l'inglese.

Se riesco, sarebbe anche utile eliminare il bisogno di usare uno script bash per eseguire l'ssh (ci sono idee?)

**Anche questo README è da sistemare**
