
if (window.location.pathname === '/checkPj') {
    const regex = /[\p{Script_Extensions=Hiragana}\p{Script_Extensions=Katakana}\wー]+/ug;
    const regex2 = /[、　]/g; // コンマと全角空白に一致する正規表現
    var existName = [];
    var existNameP = [];
    let am = isExistName(document.querySelector(".am-rest").firstChild.textContent.trim());
    let pm = isExistName(document.querySelector(".pm-rest").firstChild.textContent.trim());
    console.log(am);
    console.log(pm);
    if (am == "AM" && (pm == "PM" || pm == "試食会")) {
        amRest();
        pmRest();
    } else if (am == "AM") {
        amRest();
    } else {
        pmRest();
    }
    function amRest() {
        //AM
        let subpj = isExistName(document.querySelector(".subpj").firstChild.textContent.trim());
        let leader = isExistName(document.querySelector(".leader").firstChild.textContent.trim());
        let returnleader = isExistName(document.querySelector(".returnleader").firstChild.textContent.trim());
        let gatekeeper = isExistName(document.querySelector(".gatekeeper").firstChild.textContent.trim());
        let gatekeeper2 = isExistName(document.querySelector(".gatekeeper2").firstChild.textContent.trim());
        let cloak = isExistName(document.querySelector(".cloak").firstChild.textContent.trim());
        let cloaksub = isExistName(document.querySelector(".cloaksub").firstChild.textContent.trim());
        let coffee = isExistName(document.querySelector(".coffee").firstChild.textContent.trim());
        let drinkmain = isExistName(document.querySelector(".drinkmain").firstChild.textContent.trim());
        let drinksub = isExistName(document.querySelector(".drinksub").firstChild.textContent.trim());
        let drinksubsub = isExistName(document.querySelector(".drinksubsub").firstChild.textContent.trim());
        let champagne = isExistName(document.querySelector(".champagne").firstChild.textContent.trim());
        let silver = isExistName(document.querySelector(".silver").firstChild.textContent.trim());
        let ape = isExistName(document.querySelector(".ape").firstChild.textContent.trim());
        let toiletman = isExistName(document.querySelector(".toiletman").firstChild.textContent.trim());
        let toiletlady = isExistName(document.querySelector(".toiletlady").firstChild.textContent.trim());

        // console.log(existName);
        // rest要素内のp要素に書き込む
        let restsubpj = document.querySelector(".rest-subpj");
        restsubpj.textContent += subpj;
        let restleader = document.querySelector(".rest-leader");
        restleader.textContent += leader;
        // restleader.textContent += wash2;
        let restgatekeeper = document.querySelector(".rest-gatekeeper");
        restgatekeeper.textContent += (gatekeeper || "") + (gatekeeper && gatekeeper2 ? "," : "") + (gatekeeper2 || "");
        let restcloak = document.querySelector(".rest-cloak");
        restcloak.textContent += (cloak || "") + (cloak && (returnleader || cloaksub) ? "," : "") + (returnleader || "") + (returnleader && cloaksub ? "," : "") + (cloaksub || "");
        let restcoffee = document.querySelector(".rest-coffee");
        restcoffee.textContent += coffee;
        // restcoffee.textContent += wash1;
        let restdrink = document.querySelector(".rest-drink");
        restdrink.textContent += (drinkmain || "") + (drinkmain && (drinksub || drinksubsub) ? "," : "") + (drinksub || "") + (drinksub && drinksubsub ? "," : "") + (drinksubsub || "");
        let restchampagne = document.querySelector(".rest-champagne");
        restchampagne.textContent += (champagne || "") + (champagne && toiletlady ? "," : "") + (toiletlady || "");
        let restsilver = document.querySelector(".rest-silver");
        restsilver.textContent += silver;
        let restape = document.querySelector(".rest-ape");
        restape.textContent += (ape || "") + (ape && toiletman ? "," : "") + (toiletman || "");
    }
    function pmRest() {
        //PM
        let subpjP = isExistNameP(document.querySelector(".subpjP").firstChild.textContent.trim());
        let leaderP = isExistNameP(document.querySelector(".leaderP").firstChild.textContent.trim());
        let returnleaderP = isExistNameP(document.querySelector(".returnleaderP").firstChild.textContent.trim());
        let gatekeeperP = isExistNameP(document.querySelector(".gatekeeperP").firstChild.textContent.trim());
        let gatekeeper2P = isExistNameP(document.querySelector(".gatekeeper2P").firstChild.textContent.trim());
        let cloakP = isExistNameP(document.querySelector(".cloakP").firstChild.textContent.trim());
        let cloaksubP = isExistNameP(document.querySelector(".cloaksubP").firstChild.textContent.trim());
        let coffeeP = isExistNameP(document.querySelector(".coffeeP").firstChild.textContent.trim());
        let drinkmainP = isExistNameP(document.querySelector(".drinkmainP").firstChild.textContent.trim());
        let drinksubP = isExistNameP(document.querySelector(".drinksubP").firstChild.textContent.trim());
        let drinksubsubP = isExistNameP(document.querySelector(".drinksubsubP").firstChild.textContent.trim());
        let champagneP = isExistNameP(document.querySelector(".champagneP").firstChild.textContent.trim());
        let silverP = isExistNameP(document.querySelector(".silverP").firstChild.textContent.trim());
        let apeP = isExistNameP(document.querySelector(".apeP").firstChild.textContent.trim());
        let toiletmanP = isExistNameP(document.querySelector(".toiletmanP").firstChild.textContent.trim());
        let toiletladyP = isExistNameP(document.querySelector(".toiletladyP").firstChild.textContent.trim());
        // let wash1P = isExistName(document.querySelector(".wash1P").firstChild.textContent.trim());
        // let wash2P = isExistName(document.querySelector(".wash2P").firstChild.textContent.trim());

        let restsubpjP = document.querySelector(".rest-subpjP");
        // console.log(restsubpjP);
        restsubpjP.textContent += subpjP;
        let restleaderP = document.querySelector(".rest-leaderP");
        restleaderP.textContent += leaderP;
        // restleaderP.textContent += wash2P;
        let restgatekeeperP = document.querySelector(".rest-gatekeeperP");
        restgatekeeperP.textContent += (gatekeeperP || "") + (gatekeeperP && gatekeeper2P ? "," : "") + (gatekeeper2P || "");
        let restcloakP = document.querySelector(".rest-cloakP");
        restcloakP.textContent += (cloakP || "") + (cloakP && (returnleaderP || cloaksubP) ? "," : "") + (returnleaderP || "") + (returnleaderP && cloaksubP ? "," : "") + (cloaksubP || "");
        let restcoffeeP = document.querySelector(".rest-coffeeP");
        restcoffeeP.textContent += coffeeP;
        // restcoffeeP.textContent += wash1P;
        let restdrinkP = document.querySelector(".rest-drinkP");
        restdrinkP.textContent += (drinkmainP || "") + (drinkmainP && (drinksubP || drinksubsubP) ? "," : "") + (drinksubP || "") + (drinksubP && drinksubsubP ? "," : "") + (drinksubsubP || "");
        let restchampagneP = document.querySelector(".rest-champagneP");
        restchampagneP.textContent += (champagneP || "") + (champagneP && toiletladyP ? "," : "") + (toiletladyP || "");
        let restsilverP = document.querySelector(".rest-silverP");
        restsilverP.textContent += silverP;
        let restapeP = document.querySelector(".rest-apeP");
        restapeP.textContent += (apeP || "") + (apeP && toiletmanP ? "," : "") + (toiletmanP || "");
    }
    function isExistName(obj) {
        if (!obj || obj.length === 0) {
            return "";
        }
        // カンマとして扱われるべき文字をカンマに置換します
        let replacedObj = obj.replace(regex2, ",");
        // 入力文字列内のすべての正規表現に一致する部分文字列を取得します
        let matches = replacedObj.match(regex);
        let returnName = "";
        for (let i = 0; i < matches.length; i++) {
            // 現在の部分文字列が既に existName 配列に含まれているかどうかを確認します
            if (existName.indexOf(matches[i]) !== -1) {
                // 含まれている場合、existName 配列が空かどうかを確認します

            } else {
                // 現在の部分文字列が existName 配列に含まれていない場合、現在の部分文字列を existName 配列に追加し、戻り値に現在の部分文字列を追加します
                existName.push(matches[i]);
                returnName += matches[i] + " ";
            }
        }
        // トリムされた名前の文字列を戻します
        // console.log(returnName.trim());
        // console.log("Done!");
        return returnName.trim();
    }


    function isExistNameP(obj) {
        if (!obj || obj.length === 0) {
            return "";
        }
        // カンマとして扱われるべき文字をカンマに置換します
        let replacedObj = obj.replace(regex2, ",");
        // 入力文字列内のすべての正規表現に一致する部分文字列を取得します
        let matches = replacedObj.match(regex);
        let returnName = "";
        for (let i = 0; i < matches.length; i++) {
            // 現在の部分文字列が既に existName 配列に含まれているかどうかを確認します
            if (existNameP.indexOf(matches[i]) !== -1) {
                // 含まれている場合、existName 配列が空かどうかを確認します
            } else {
                // 現在の部分文字列が existName 配列に含まれていない場合、現在の部分文字列を existName 配列に追加し、戻り値に現在の部分文字列を追加します
                existNameP.push(matches[i]);
                returnName += matches[i] + " ";
            }
        }
        // トリムされた名前の文字列を戻します
        // console.log(returnName.trim());
        return returnName.trim();
    }
}