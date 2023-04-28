// towel2Pの文字列を取得する
const subpj = document.querySelector(".subpj").firstChild.textContent.trim();
const leader = document.querySelector(".leader").firstChild.textContent.trim();
const returnleader = document.querySelector(".returnleader").firstChild.textContent.trim();
const gatekeeper = document.querySelector(".gatekeeper").firstChild.textContent.trim();
const parkkeeper1 = document.querySelector(".parkkeeper1").firstChild.textContent.trim();
const parkkeeper2 = document.querySelector(".parkkeeper2").firstChild.textContent.trim();
const cloak = document.querySelector(".cloak").firstChild.textContent.trim();
const cloaksub = document.querySelector(".cloaksub").firstChild.textContent.trim();
const coffee = document.querySelector(".coffee").firstChild.textContent.trim();
const drinkmain = document.querySelector(".drinkmain").firstChild.textContent.trim();
const drinksub = document.querySelector(".drinksub").firstChild.textContent.trim();
const drinksubsub = document.querySelector(".drinksubsub").firstChild.textContent.trim();
const champagne1 = document.querySelector(".champagne1").firstChild.textContent.trim();
const champagne2 = document.querySelector(".champagne2").firstChild.textContent.trim();
const silver1 = document.querySelector(".silver1").firstChild.textContent.trim();
const silver2 = document.querySelector(".silver2").firstChild.textContent.trim();
const ape1 = document.querySelector(".ape1").firstChild.textContent.trim();
const ape2 = document.querySelector(".ape2").firstChild.textContent.trim();
const toiletman = document.querySelector(".toiletman").firstChild.textContent.trim();
const toiletlady = document.querySelector(".toiletlady").firstChild.textContent.trim();
// const wash1 = document.querySelector(".wash1").firstChild.textContent.trim();
// const wash2 = document.querySelector(".wash2").firstChild.textContent.trim();


// rest要素内のp要素に書き込む
const restsubpj = document.querySelector(".rest-subpj");
restsubpj.textContent += subpj;
const restleader = document.querySelector(".rest-leader");
restleader.textContent += leader;
// restleader.textContent += wash2;
const restgatekeeper = document.querySelector(".rest-gatekeeper");
restgatekeeper.textContent += (gatekeeper || "") + (gatekeeper && (parkkeeper1 || parkkeeper2) ? "," : "") + (parkkeeper1 || "") + (parkkeeper1 && parkkeeper2 ? "," : "") + (parkkeeper2 || "");
const restcloak = document.querySelector(".rest-cloak");
restcloak.textContent += (cloak || "") + (cloak && (returnleader || cloaksub) ? "," : "") + (returnleader || "") + (returnleader && cloaksub ? "," : "") + (cloaksub || "");
const restcoffee = document.querySelector(".rest-coffee");
restcoffee.textContent += coffee;
// restcoffee.textContent += wash1;
const restdrink = document.querySelector(".rest-drink");
restdrink.textContent += (drinkmain || "") + (drinkmain && (drinksub || drinksubsub) ? "," : "") + (drinksub || "") + (drinksub && drinksubsub ? "," : "") + (drinksubsub || "");
const restchampagne = document.querySelector(".rest-champagne");
restchampagne.textContent += (champagne1 || "") + (champagne1 && (champagne2 || toiletlady) ? "," : "") + (champagne2 || "") + (champagne2 && toiletlady ? "," : "") + (toiletlady || "");
const restsilver = document.querySelector(".rest-silver");
restsilver.textContent += (silver1 || "") + (silver1 && silver2 ? "," : "") + (silver2 || "");
const restape = document.querySelector(".rest-ape");
restape.textContent += (ape1 || "") + (ape1 && (ape2 || toiletman) ? "," : "") + (ape2 || "") + (ape2 && toiletman ? "," : "") + (toiletman || "");


//PM

const subpjP = document.querySelector(".subpjP").firstChild.textContent.trim();
const leaderP = document.querySelector(".leaderP").firstChild.textContent.trim();
const returnleaderP = document.querySelector(".returnleaderP").firstChild.textContent.trim();
const gatekeeperP = document.querySelector(".gatekeeperP").firstChild.textContent.trim();
const parkkeeper1P = document.querySelector(".parkkeeper1P").firstChild.textContent.trim();
const parkkeeper2P = document.querySelector(".parkkeeper2P").firstChild.textContent.trim();
const cloakP = document.querySelector(".cloakP").firstChild.textContent.trim();
const cloaksubP = document.querySelector(".cloaksubP").firstChild.textContent.trim();
const coffeeP = document.querySelector(".coffeeP").firstChild.textContent.trim();
const drinkmainP = document.querySelector(".drinkmainP").firstChild.textContent.trim();
const drinksubP = document.querySelector(".drinksubP").firstChild.textContent.trim();
const drinksubsubP = document.querySelector(".drinksubsubP").firstChild.textContent.trim();
const champagne1P = document.querySelector(".champagne1P").firstChild.textContent.trim();
const champagne2P = document.querySelector(".champagne2P").firstChild.textContent.trim();
const silver1P = document.querySelector(".silver1P").firstChild.textContent.trim();
const silver2P = document.querySelector(".silver2P").firstChild.textContent.trim();
const ape1P = document.querySelector(".ape1P").firstChild.textContent.trim();
const ape2P = document.querySelector(".ape2P").firstChild.textContent.trim();
const toiletmanP = document.querySelector(".toiletmanP").firstChild.textContent.trim();
const toiletladyP = document.querySelector(".toiletladyP").firstChild.textContent.trim();
// const wash1P = document.querySelector(".wash1P").firstChild.textContent.trim();
// const wash2P = document.querySelector(".wash2P").firstChild.textContent.trim();

const restsubpjP = document.querySelector(".rest-subpjP");
restsubpjP.textContent += subpjP;
const restleaderP = document.querySelector(".rest-leaderP");
restleaderP.textContent += leaderP;
// restleaderP.textContent += wash2P;
const restgatekeeperP = document.querySelector(".rest-gatekeeperP");
restgatekeeperP.textContent += (gatekeeperP || "") + (gatekeeperP && (parkkeeper1P || parkkeeper2P) ? "," : "") + (parkkeeper1P || "") + (parkkeeper1P && parkkeeper2P ? "," : "") + (parkkeeper2P || "");
const restcloakP = document.querySelector(".rest-cloakP");
restcloakP.textContent += (cloakP || "") + (cloakP && (returnleaderP || cloaksubP) ? "," : "") + (returnleaderP || "") + (returnleaderP && cloaksubP ? "," : "") + (cloaksubP || "");
const restcoffeeP = document.querySelector(".rest-coffeeP");
restcoffeeP.textContent += coffeeP;
// restcoffeeP.textContent += wash1P;
const restdrinkP = document.querySelector(".rest-drinkP");
restdrinkP.textContent += (drinkmainP || "") + (drinkmainP && (drinksubP || drinksubsubP) ? "," : "") + (drinksubP || "") + (drinksubP && drinksubsubP ? "," : "") + (drinksubsubP || "");
const restchampagneP = document.querySelector(".rest-champagneP");
restchampagneP.textContent += (champagne1P || "") + (champagne1P && (champagne2P || toiletladyP) ? "," : "") + (champagne2P || "") + (champagne2P && toiletladyP ? "," : "") + (toiletladyP || "");
const restsilverP = document.querySelector(".rest-silverP");
restsilverP.textContent += (silver1P || "") + (silver1P && silver2P ? "," : "") + (silver2P || "");
const restapeP = document.querySelector(".rest-apeP");
restapeP.textContent += (ape1P || "") + (ape1P && (ape2P || toiletmanP) ? "," : "") + (ape2P || "") + (ape2P && toiletmanP ? "," : "") + (toiletmanP || "");

