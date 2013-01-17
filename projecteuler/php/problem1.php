<?php
$sum = 0;
for ($i = 1; $i < 1000; $i++)
    $sum+= (($i % 3 === 0) or ($i % 5 === 0)) ? $i : 0;
 
echo $sum;
