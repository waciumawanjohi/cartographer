var ge=Object.defineProperty;var X=Object.getOwnPropertySymbols;var he=Object.prototype.hasOwnProperty,ye=Object.prototype.propertyIsEnumerable;var J=(t,e,r)=>e in t?ge(t,e,{enumerable:!0,configurable:!0,writable:!0,value:r}):t[e]=r,W=(t,e)=>{for(var r in e||(e={}))he.call(e,r)&&J(t,r,e[r]);if(X)for(var r of X(e))ye.call(e,r)&&J(t,r,e[r]);return t};import{S as M,i as P,s as R,c as Q,e as d,a as I,b as f,d as U,f as h,g as _,u as ee,h as te,j as re,t as C,k as T,l as y,o as B,m as V,n as x,r as ne,p as be,q as N,v as K,w as Z,x as se,y as ve,z as we,A as ke,B as A,C as E,D as F,E as O,F as _e,G as Ce,H as G,I as xe,J as Se,U as oe,K as S,L as ae,M as $e,N as Te,O as Ie,P as Le,Q as Me,R as Pe,T as w,V as ie,W as Re,X as De,Y as b,Z as v,_ as q,$ as H,a0 as Ae,a1 as Ee,a2 as Fe}from"./vendor.95398b96.js";const je=function(){const e=document.createElement("link").relList;if(e&&e.supports&&e.supports("modulepreload"))return;for(const s of document.querySelectorAll('link[rel="modulepreload"]'))n(s);new MutationObserver(s=>{for(const o of s)if(o.type==="childList")for(const a of o.addedNodes)a.tagName==="LINK"&&a.rel==="modulepreload"&&n(a)}).observe(document,{childList:!0,subtree:!0});function r(s){const o={};return s.integrity&&(o.integrity=s.integrity),s.referrerpolicy&&(o.referrerPolicy=s.referrerpolicy),s.crossorigin==="use-credentials"?o.credentials="include":s.crossorigin==="anonymous"?o.credentials="omit":o.credentials="same-origin",o}function n(s){if(s.ep)return;s.ep=!0;const o=r(s);fetch(s.href,o)}};je();const ze=t=>({}),ce=t=>({}),Ne=t=>({}),le=t=>({});function Ke(t){let e;return{c(){e=d("div"),e.innerHTML="<em>missing &lt;slot=&quot;left&quot;&gt;</em>",U(e,"text-align","center")},m(r,n){h(r,e,n)},d(r){r&&y(e)}}}function Ue(t){let e;return{c(){e=d("em"),e.textContent='missing <slot="right">'},m(r,n){h(r,e,n)},d(r){r&&y(e)}}}function Oe(t){let e,r,n,s,o,a,i,c;const l=t[4].left,m=Q(l,t,t[3],le),u=m||Ke(),p=t[4].right,$=Q(p,t,t[3],ce),L=$||Ue();return{c(){e=d("div"),r=d("div"),u&&u.c(),n=I(),s=d("div"),o=I(),a=d("div"),i=d("div"),L&&L.c(),f(r,"class","hidden md:flex flex-col"),U(r,"width",t[2]),f(s,"id","resizeHandler"),f(s,"class","hover:bg-sky-400 bg-sky-700 cursor-col-resize pr-1"),U(i,"text-align","center"),f(a,"class","flex-1 flex flex-col overflow-hidden"),f(e,"class","flex-1 flex overflow-hidden")},m(g,D){h(g,e,D),_(e,r),u&&u.m(r,null),t[5](r),_(e,n),_(e,s),t[6](s),_(e,o),_(e,a),_(a,i),L&&L.m(i,null),c=!0},p(g,[D]){m&&m.p&&(!c||D&8)&&ee(m,l,g,g[3],c?re(l,g[3],D,Ne):te(g[3]),le),(!c||D&4)&&U(r,"width",g[2]),$&&$.p&&(!c||D&8)&&ee($,p,g,g[3],c?re(p,g[3],D,ze):te(g[3]),ce)},i(g){c||(C(u,g),C(L,g),c=!0)},o(g){T(u,g),T(L,g),c=!1},d(g){g&&y(e),u&&u.d(g),t[5](null),t[6](null),L&&L.d(g)}}}function qe(t,e,r){let{$$slots:n={},$$scope:s}=e,o,a,i="50%";const c=p=>{const $=p.pageX-a.getBoundingClientRect().left;$>50&&r(2,i=`${$}px`)},l=()=>{window.removeEventListener("mousemove",c)};B(()=>{o.addEventListener("mousedown",p=>{p.preventDefault(),window.addEventListener("mousemove",c),window.addEventListener("mouseup",l)})});function m(p){V[p?"unshift":"push"](()=>{a=p,r(1,a)})}function u(p){V[p?"unshift":"push"](()=>{o=p,r(0,o)})}return t.$$set=p=>{"$$scope"in p&&r(3,s=p.$$scope)},[o,a,i,s,n,m,u]}class We extends M{constructor(e){super();P(this,e,qe,Oe,R,{})}}function Ve(t){let e,r;return{c(){e=d("pre"),f(e,"class",r=t[1].class)},m(n,s){h(n,e,s),t[3](e)},p(n,[s]){s&2&&r!==(r=n[1].class)&&f(e,"class",r)},i:x,o:x,d(n){n&&y(e),t[3](null)}}}function He(t,e,r){let{doc:n=""}=e,s=null;ne.initialize({startOnLoad:!1,theme:"forest"});const o=()=>{try{ne.mermaidAPI.render("graph-div",n,i=>{r(0,s.innerHTML=i,s)})}catch{}};B(()=>{o()}),be(()=>{o()});function a(i){V[i?"unshift":"push"](()=>{s=i,r(0,s)})}return t.$$set=i=>{r(1,e=N(N({},e),K(i))),"doc"in i&&r(2,n=i.doc)},e=K(e),[s,e,n,a]}class Be extends M{constructor(e){super();P(this,e,He,Ve,R,{doc:2})}}const Ze=t=>{let e=["flowchart RL","classDef not-found fill:#f66;"],r={};return t.spec.resources&&t.spec.resources.forEach(n=>{if(!n.name)return;let s=`res-${n.name}`,o=`${n.name}`;r[s]=o,e.push(`${s}["${o}"]`),["sources","images","configs"].forEach((a,i)=>{n[a]&&n[a].forEach(c=>{if(r[`res-${c.resource}`])e.push(`${s} --> res-${c.resource}`);else{let l=`not-found-${c.resource}-${i}`;e.push(`${s} --> ${l}`),e.push(`${l}["not-found"]`),e.push(`class ${l} not-found`)}})})}),e.join(`
`)},j=ke(`---
apiVersion: carto.run/v1alpha1
kind: ClusterSupplyChain
metadata:
  name: supply-chain
spec:
  selector:
    app.tanzu.vmware.com/workload-type: web

  resources:
    - name: source-provider
      templateRef:
        kind: ClusterSourceTemplate
        name: source

    - name: image-builder
      templateRef:
        kind: ClusterImageTemplate
        name: image
      params:
        - name: image_prefix
          value: "pref-"
      sources:
        - resource: source-provider
          name: source

    - name: config-provider
      templateRef:
        kind: ClusterConfigTemplate
        name: app-config
      images:
        - resource: image-builder
          name: image

    - name: git-writer
      templateRef:
        kind: ClusterTemplate
        name: git-writer
      configs:
        - resource: config-provider
          name: data
`),Ge=Z(j,t=>{try{return se(t)}catch(e){console.log(`could not parse to yaml object: ${e}`)}}),Ye=Z(j,t=>{try{let e=new TextEncoder().encode(t),r=ve(e,{options:9});return we(r,!0)}catch(e){console.log(`could not compress document: ${e}`)}}),Xe=Z(Ge,(t,e)=>{try{e(Ze(t))}catch(r){console.log(`could not parse to mermaid: ${r}`)}});function Je(t){let e,r;return e=new Be({props:{doc:t[0],class:t[1].class}}),{c(){A(e.$$.fragment)},m(n,s){E(e,n,s),r=!0},p(n,[s]){const o={};s&1&&(o.doc=n[0]),s&2&&(o.class=n[1].class),e.$set(o)},i(n){r||(C(e.$$.fragment,n),r=!0)},o(n){T(e.$$.fragment,n),r=!1},d(n){F(e,n)}}}function Qe(t,e,r){let n;return O(t,Xe,s=>r(0,n=s)),t.$$set=s=>{r(1,e=N(N({},e),K(s)))},e=K(e),[n,e]}class et extends M{constructor(e){super();P(this,e,Qe,Je,R,{})}}function tt(t){let e,r,n,s,o,a,i;return{c(){e=d("h1"),e.innerHTML='<span class="uppercase">Cartographer Live Editor Help</span>',r=I(),n=d("p"),n.innerHTML=`This editor is a spike to gauge interest and investigate the possibilities.

    Provide feedback by commenting on the
    <a href="https://github.com/vmware-tanzu/cartographer/issues/566" target="_blank" class="underline text-blue-700 hover:text-orange-700">original issue</a>, but please make sure you don&#39;t report
    something already captured.`,s=I(),o=d("h2"),o.textContent="Shortcuts",a=I(),i=d("div"),i.innerHTML=`<div class="text-right">Autocomplete</div> 
    <div><kbd class="svelte-1t0r1ka">Ctrl</kbd> + <kbd class="svelte-1t0r1ka">Space</kbd></div> 
    <div class="text-right">Show info</div> 
    <div><kbd class="svelte-1t0r1ka">Cmd</kbd> + <kbd class="svelte-1t0r1ka">K</kbd>, <kbd class="svelte-1t0r1ka">Cmd</kbd> + <kbd class="svelte-1t0r1ka">I</kbd></div> 
    <div class="text-right">Command Palette</div> 
    <div><kbd class="svelte-1t0r1ka">F1</kbd></div>`,f(e,"class","text-md pb-2"),f(n,"class","text-sm pt-1"),f(o,"class","text-md pt-3 pb-2"),f(i,"class","grid grid-cols-2 gap-x-3 gap-y-3 w-full text-sm")},m(c,l){h(c,e,l),h(c,r,l),h(c,n,l),h(c,s,l),h(c,o,l),h(c,a,l),h(c,i,l)},p:x,i:x,o:x,d(c){c&&y(e),c&&y(r),c&&y(n),c&&y(s),c&&y(o),c&&y(a),c&&y(i)}}}function rt(t){return navigator.platform.toUpperCase().indexOf("MAC")>=0,[]}class nt extends M{constructor(e){super();P(this,e,rt,tt,R,{})}}function st(t){let e,r,n,s,o;return{c(){e=d("h1"),e.innerHTML='<span class="uppercase">Link copied to clipboard!</span>',r=I(),n=d("form"),s=d("input"),o=_e(`
    was copied to your clipboard!`),f(e,"class","text-md pb-2"),f(s,"class","p-2 border-2 border-sky-700 w-full"),f(s,"type","text"),s.value=window.location.href},m(a,i){h(a,e,i),h(a,r,i),h(a,n,i),_(n,s),_(n,o)},p:x,i:x,o:x,d(a){a&&y(e),a&&y(r),a&&y(n)}}}class ot extends M{constructor(e){super();P(this,e,null,st,R,{})}}function at(t){let e,r,n;return{c(){e=d("button"),e.textContent="Share",f(e,"class","text-lg no-underline text-grey-darkest hover:text-orange-300 ml-2")},m(s,o){h(s,e,o),r||(n=G(e,"click",function(){Se(t[7])&&t[7].apply(this,arguments)}),r=!0)},p(s,o){t=s},d(s){s&&y(e),r=!1,n()}}}function it(t){let e,r,n,s,o,a,i,c,l,m;return o=new Ce({props:{text:t[0],$$slots:{default:[at,({copy:u})=>({7:u}),({copy:u})=>u?128:0]},$$scope:{ctx:t}}}),o.$on("copy",t[3]),o.$on("fail",t[4]),{c(){e=d("nav"),r=d("div"),r.innerHTML='<h1 class="text-lg uppercase">Cartographer Live Editor</h1>',n=I(),s=d("div"),A(o.$$.fragment),a=I(),i=d("button"),i.textContent="Help",f(r,"class","mb-2 sm:mb-0"),f(i,"class","text-lg no-underline text-grey-darkest hover:text-orange-300 ml-2"),f(e,"class","font-sans flex flex-col text-center sm:flex-row sm:text-left sm:justify-between py-2 px-6 bg-sky-700 text-sky-50 shadow sm:items-baseline w-full")},m(u,p){h(u,e,p),_(e,r),_(e,n),_(e,s),E(o,s,null),_(s,a),_(s,i),c=!0,l||(m=G(i,"click",t[5]),l=!0)},p(u,[p]){const $={};p&1&&($.text=u[0]),p&384&&($.$$scope={dirty:p,ctx:u}),o.$set($)},i(u){c||(C(o.$$.fragment,u),c=!0)},o(u){T(o.$$.fragment,u),c=!1},d(u){u&&y(e),F(o),l=!1,m()}}}function ct(t,e,r){let n;O(t,j,m=>r(2,n=m));const{open:s}=xe("simple-modal");let o;const a=()=>window.location.href,i=()=>s(ot),c=()=>alert("Something went wrong with the copy to clipboard, sorry!"),l=()=>s(nt);return t.$$.update=()=>{t.$$.dirty&4&&r(0,o=a())},[o,s,n,i,c,l]}class lt extends M{constructor(e){super();P(this,e,ct,it,R,{})}}function ut(t){let e,r;return{c(){e=d("div"),f(e,"class",r="monaco-editor-container "+t[1].class+" svelte-oc3rgm")},m(n,s){h(n,e,s),t[2](e)},p(n,[s]){s&2&&r!==(r="monaco-editor-container "+n[1].class+" svelte-oc3rgm")&&f(e,"class",r)},i:x,o:x,d(n){n&&y(e),t[2](null)}}}function dt(t,e,r){let n;O(t,j,l=>r(4,n=l));let s,o;const a=oe.parse("https://cartographer.sh/file.yaml"),i=S.createModel("","yaml",a);B(()=>{i.setValue(n),s=S.create(o,{automaticLayout:!0,model:i}),s.onDidChangeModelContent(l=>{ae(j,n=s.getValue(),n)})});function c(l){V[l?"unshift":"push"](()=>{o=l,r(0,o)})}return t.$$set=l=>{r(1,e=N(N({},e),K(l)))},e=K(e),[o,e,c]}class mt extends M{constructor(e){super();P(this,e,dt,ut,R,{})}}const{window:pt}=Te;function ue(t){let e,r;return e=new We({props:{$$slots:{right:[gt],left:[ft]},$$scope:{ctx:t}}}),{c(){A(e.$$.fragment)},m(n,s){E(e,n,s),r=!0},i(n){r||(C(e.$$.fragment,n),r=!0)},o(n){T(e.$$.fragment,n),r=!1},d(n){F(e,n)}}}function ft(t){let e,r;return e=new mt({props:{slot:"left",class:"h-full m-2"}}),{c(){A(e.$$.fragment)},m(n,s){E(e,n,s),r=!0},p:x,i(n){r||(C(e.$$.fragment,n),r=!0)},o(n){T(e.$$.fragment,n),r=!1},d(n){F(e,n)}}}function gt(t){let e,r;return e=new et({props:{slot:"right",class:"content-center"}}),{c(){A(e.$$.fragment)},m(n,s){E(e,n,s),r=!0},p:x,i(n){r||(C(e.$$.fragment,n),r=!0)},o(n){T(e.$$.fragment,n),r=!1},d(n){F(e,n)}}}function ht(t){let e,r,n,s;e=new lt({});let o=t[0]&&ue(t);return{c(){A(e.$$.fragment),r=I(),n=d("div"),o&&o.c(),f(n,"class","h-full flex flex-col overflow-hidden")},m(a,i){E(e,a,i),h(a,r,i),h(a,n,i),o&&o.m(n,null),s=!0},p(a,i){a[0]?o?i&1&&C(o,1):(o=ue(a),o.c(),C(o,1),o.m(n,null)):o&&(Ie(),T(o,1,1,()=>{o=null}),Le())},i(a){s||(C(e.$$.fragment,a),C(o),s=!0)},o(a){T(e.$$.fragment,a),T(o),s=!1},d(a){F(e,a),a&&y(r),a&&y(n),o&&o.d()}}}function yt(t){let e,r,n,s,o;return r=new $e({props:{unstyled:!0,closeButton:!1,classBg:"fixed top-0 left-0 w-screen h-screen flex flex-col justify-center bg-gray-400/[.6] z-50",classWindowWrap:"relative m-2 max-h-full",classWindow:"relative w-1/3 max-w-full max-h-full my-2 mx-auto text-black border border-sky-600 shadow-lg bg-white",classContent:"relative p-2 overflow-auto",$$slots:{default:[ht]},$$scope:{ctx:t}}}),{c(){e=d("main"),A(r.$$.fragment),f(e,"class","text-primary-content"),U(e,"height","95vh")},m(a,i){h(a,e,i),E(r,e,null),n=!0,s||(o=G(pt,"load",t[1]),s=!0)},p(a,[i]){const c={};i&33&&(c.$$scope={dirty:i,ctx:a}),r.$set(c)},i(a){n||(C(r.$$.fragment,a),n=!0)},o(a){T(r.$$.fragment,a),n=!1},d(a){a&&y(e),F(r),s=!1,o()}}}function bt(t,e,r){let n,s;O(t,j,c=>r(3,n=c)),O(t,Ye,c=>r(2,s=c));let o=!1;const a=(c,l)=>{if(!c||!l)return;let m=new URL(window.location.href);m.searchParams.set("pako",l),window.history.pushState("","",m)},i=()=>{let l=new URL(window.location.href).searchParams.get("pako");l&&ae(j,n=Me(Pe(l),{to:"string"}),n),r(0,o=!0)};return t.$$.update=()=>{t.$$.dirty&5&&a(o,s)},[o,i,s]}class vt extends M{constructor(e){super();P(this,e,bt,yt,R,{})}}var k="yaml";function wt(t){switch(t){case H.Error:return q.Error;case H.Warning:return q.Warning;case H.Information:return q.Info;case H.Hint:return q.Hint;default:return q.Info}}function kt(t){return{severity:wt(t.severity),startLineNumber:t.range.start.line+1,startColumn:t.range.start.character+1,endLineNumber:t.range.end.line+1,endColumn:t.range.end.character+1,message:t.message,code:String(t.code),source:t.source}}function _t(t,e){const r=new Map,n=async i=>{(await t()).resetSchema(String(i))},s=async i=>{const m=(await(await t(i)).doValidation(String(i))).map(kt),u=S.getModel(i);u&&u.getLanguageId()===k&&S.setModelMarkers(u,k,m)},o=i=>{if(i.getLanguageId()!==k)return;let c;r.set(String(i.uri),i.onDidChangeContent(()=>{clearTimeout(c),c=setTimeout(()=>s(i.uri),500)})),s(i.uri)},a=i=>{S.setModelMarkers(i,k,[]);const c=String(i.uri),l=r.get(c);l&&(l.dispose(),r.delete(c))};S.onDidCreateModel(o),S.onWillDisposeModel(i=>{a(i),n(i.uri)}),S.onDidChangeModelLanguage(i=>{a(i.model),o(i.model),n(i.model.uri)}),e.onDidChange(()=>{for(const i of S.getModels())i.getLanguageId()===k&&(a(i),o(i))});for(const i of S.getModels())o(i)}function Y(t){if(!!t)return{character:t.column-1,line:t.lineNumber-1}}function z(t){if(!!t)return new ie(t.start.line+1,t.start.character+1,t.end.line+1,t.end.character+1)}function Ct(t){const e=w.CompletionItemKind;switch(t){case b.Text:return e.Text;case b.Method:return e.Method;case b.Function:return e.Function;case b.Constructor:return e.Constructor;case b.Field:return e.Field;case b.Variable:return e.Variable;case b.Class:return e.Class;case b.Interface:return e.Interface;case b.Module:return e.Module;case b.Property:return e.Property;case b.Unit:return e.Unit;case b.Value:return e.Value;case b.Enum:return e.Enum;case b.Keyword:return e.Keyword;case b.Snippet:return e.Snippet;case b.Color:return e.Color;case b.File:return e.File;case b.Reference:return e.Reference;default:return e.Property}}function de(t){if(!!t)return{range:z(t.range),text:t.newText}}function xt(t){return{triggerCharacters:[" ",":"],async provideCompletionItems(e,r){const n=e.uri,o=await(await t(n)).doComplete(String(n),Y(r));if(!o)return;const a=e.getWordUntilPosition(r),i=new ie(r.lineNumber,a.startColumn,r.lineNumber,a.endColumn),c=o.items.map(l=>{const m={label:l.label,insertText:l.insertText||l.label,sortText:l.sortText,filterText:l.filterText,documentation:l.documentation,detail:l.detail,kind:Ct(l.kind),range:i};return l.textEdit&&(m.range=z("range"in l.textEdit?l.textEdit.range:l.textEdit.replace),m.insertText=l.textEdit.newText),l.additionalTextEdits&&(m.additionalTextEdits=l.additionalTextEdits.map(de)),l.insertTextFormat===Re.Snippet&&(m.insertTextRules=w.CompletionItemInsertTextRule.InsertAsSnippet),m});return{incomplete:o.isIncomplete,suggestions:c}}}}function St(t){return{async provideDefinition(e,r){const n=e.uri,o=await(await t(n)).doDefinition(String(n),Y(r));return o==null?void 0:o.map(a=>({originSelectionRange:a.originSelectionRange,range:z(a.targetRange),targetSelectionRange:a.targetSelectionRange,uri:oe.parse(a.targetUri)}))}}}function $t(t){return{async provideHover(e,r){const n=e.uri,o=await(await t(n)).doHover(String(n),Y(r));if(!!o)return{range:z(o.range),contents:[{value:o.contents.value}]}}}}function Tt(t){const e=w.SymbolKind;switch(t){case v.File:return e.Array;case v.Module:return e.Module;case v.Namespace:return e.Namespace;case v.Package:return e.Package;case v.Class:return e.Class;case v.Method:return e.Method;case v.Property:return e.Property;case v.Field:return e.Field;case v.Constructor:return e.Constructor;case v.Enum:return e.Enum;case v.Interface:return e.Interface;case v.Function:return e.Function;case v.Variable:return e.Variable;case v.Constant:return e.Constant;case v.String:return e.String;case v.Number:return e.Number;case v.Boolean:return e.Boolean;case v.Array:return e.Array;default:return e.Function}}function me(t){return{detail:t.detail||"",range:z(t.range),name:t.name,kind:Tt(t.kind),selectionRange:z(t.selectionRange),children:t.children.map(me),tags:[]}}function It(t){return{async provideDocumentSymbols(e){const r=e.uri,s=await(await t(r)).findDocumentSymbols(String(r));if(!!s)return s.map(me)}}}function Lt(t){return W({tabSize:t.tabSize,insertSpaces:t.insertSpaces},t)}function Mt(t){return{async provideDocumentFormattingEdits(e,r){const n=e.uri,o=await(await t(n)).format(String(n),Lt(r));if(!(!o||o.length===0))return o.map(de)}}}function Pt(t){return{range:z(t.range),tooltip:t.tooltip,url:t.target}}function Rt(t){return{async provideLinks(e){const r=e.uri;return{links:(await(await t(r)).findLinks(String(r))).map(Pt)}}}}var Dt=2*60*1e3;function At(t){let e,r,n=0;const s=()=>{e&&(e.dispose(),e=void 0),r=void 0};setInterval(()=>{if(!e)return;Date.now()-n>Dt&&s()},30*1e3),t.onDidChange(()=>s());const o=()=>(n=Date.now(),r||(e=S.createWebWorker({moduleId:"vs/language/yaml/yamlWorker",label:t.languageId,createData:{languageSettings:t.diagnosticsOptions,enableSchemaRequest:t.diagnosticsOptions.enableSchemaRequest,isKubernetes:t.diagnosticsOptions.isKubernetes,customTags:t.diagnosticsOptions.customTags}}),r=e.getProxy()),r);return async(...a)=>{const i=await o();return await e.withSyncedResources(a),i}}var Et={comments:{lineComment:"#"},brackets:[["{","}"],["[","]"],["(",")"]],autoClosingPairs:[{open:"{",close:"}"},{open:"[",close:"]"},{open:"(",close:")"},{open:'"',close:'"'},{open:"'",close:"'"}],surroundingPairs:[{open:"{",close:"}"},{open:"[",close:"]"},{open:"(",close:")"},{open:'"',close:'"'},{open:"'",close:"'"}],onEnterRules:[{beforeText:/:\s*$/,action:{indentAction:w.IndentAction.Indent}}]};function Ft(t){const e=At(t);w.registerCompletionItemProvider(k,xt(e)),w.registerHoverProvider(k,$t(e)),w.registerDefinitionProvider(k,St(e)),w.registerDocumentSymbolProvider(k,It(e)),w.registerDocumentFormattingEditProvider(k,Mt(e)),w.registerLinkProvider(k,Rt(e)),_t(e,t),w.setLanguageConfiguration(k,Et)}var pe={completion:!0,customTags:[],enableSchemaRequest:!1,format:!0,isKubernetes:!1,hover:!0,schemas:[],validate:!0,yamlVersion:"1.2"};function jt(t){const e=new De;let r=t;const n={get onDidChange(){return e.event},get languageId(){return k},get diagnosticsOptions(){return r},setDiagnosticsOptions(s){r=W(W({},pe),s),e.fire(n)}};return n}var fe=jt(pe);function zt(){return{yamlDefaults:fe}}w.yaml=zt();w.register({id:k,extensions:[".yaml",".yml"],aliases:["YAML","yaml","YML","yml"],mimetypes:["application/x-yaml"]});w.onLanguage("yaml",()=>{Ft(fe)});function Nt(t={}){w.yaml.yamlDefaults.setDiagnosticsOptions(t)}function Kt(){return new Worker("/live-editor/assets/yaml.worker.241bdf2d.js",{type:"module"})}const Ut={properties:{apiVersion:{description:"APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",type:"string"},kind:{description:"Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",type:"string"},metadata:{type:"object"},spec:{description:"Spec describes the suppply chain. More info: https://cartographer.sh/docs/latest/reference/workload/#clustersupplychain",properties:{params:{description:"Additional parameters. See: https://cartographer.sh/docs/latest/architecture/#parameter-hierarchy",items:{properties:{default:{description:"DefaultValue of the parameter. Causes the parameter to be optional; If the Owner does not specify this parameter, this value is used.","x-kubernetes-preserve-unknown-fields":!0},name:{description:"Name of the parameter. Should match a template parameter name.",type:"string"},value:{description:"Value of the parameter. If specified, owner properties are ignored.","x-kubernetes-preserve-unknown-fields":!0}},required:["name"],type:"object"},type:"array"},resources:{description:"Resources that are responsible for bringing the application to a deliverable state.",items:{properties:{configs:{description:`Configs is a list of references to other 'config' resources in this list. A config resource has the kind ClusterConfigTemplate 
 In a template, configs can be consumed as: $(configs.<name>.config)$ 
 If there is only one image, it can be consumed as: $(config)$`,items:{properties:{name:{type:"string"},resource:{type:"string"}},required:["name","resource"],type:"object"},type:"array"},images:{description:`Images is a list of references to other 'image' resources in this list. An image resource has the kind ClusterImageTemplate 
 In a template, images can be consumed as: $(images.<name>.image)$ 
 If there is only one image, it can be consumed as: $(image)$`,items:{properties:{name:{type:"string"},resource:{type:"string"}},required:["name","resource"],type:"object"},type:"array"},name:{description:"Name of the resource. Used as a reference for inputs, as well as being the name presented in workload statuses to identify this resource.",type:"string"},params:{description:`Params are a list of parameters to provide to the template in TemplateRef Template params do not have to be specified here, unless you want to force a particular value, or add a default value. 
 Parameters are consumed in a template with the syntax: $(params.<name>)$`,items:{properties:{default:{description:"DefaultValue of the parameter. Causes the parameter to be optional; If the Owner does not specify this parameter, this value is used.","x-kubernetes-preserve-unknown-fields":!0},name:{description:"Name of the parameter. Should match a template parameter name.",type:"string"},value:{description:"Value of the parameter. If specified, owner properties are ignored.","x-kubernetes-preserve-unknown-fields":!0}},required:["name"],type:"object"},type:"array"},sources:{description:`Sources is a list of references to other 'source' resources in this list. A source resource has the kind ClusterSourceTemplate 
 In a template, sources can be consumed as: $(sources.<name>.url)$ and $(sources.<name>.revision)$ 
 If there is only one source, it can be consumed as: $(source.url)$ and $(source.revision)$`,items:{properties:{name:{type:"string"},resource:{type:"string"}},required:["name","resource"],type:"object"},type:"array"},templateRef:{description:"TemplateRef identifies the template used to produce this resource",properties:{kind:{description:"Kind of the template to apply",enum:["ClusterSourceTemplate","ClusterImageTemplate","ClusterTemplate","ClusterConfigTemplate"],type:"string"},name:{description:"Name of the template to apply",minLength:1,type:"string"}},required:["kind","name"],type:"object"}},required:["name","templateRef"],type:"object"},type:"array"},selector:{additionalProperties:{type:"string"},description:"Specifies the label key-value pairs used to select workloads See: https://cartographer.sh/docs/v0.1.0/architecture/#selectors",type:"object"},serviceAccountRef:{description:`ServiceAccountName refers to the Service account with permissions to create resources submitted by the supply chain. 
 If not set, Cartographer will use serviceAccountName from supply chain. 
 If that is also not set, Cartographer will use the default service account in the workload's namespace.`,properties:{name:{description:"Name of the service account being referred to",type:"string"},namespace:{description:"Namespace of the service account being referred to if omitted, the Owner's namespace is used.",type:"string"}},required:["name"],type:"object"}},required:["resources","selector"],type:"object"},status:{description:"Status conforms to the Kubernetes conventions: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties",properties:{conditions:{items:{description:'Condition contains details for one aspect of the current state of this API Resource. --- This struct is intended for direct use as an array at the field path .status.conditions.  For example, type FooStatus struct{ // Represents the observations of a foo\'s current state. // Known .status.conditions.type are: "Available", "Progressing", and "Degraded" // +patchMergeKey=type // +patchStrategy=merge // +listType=map // +listMapKey=type Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"` \n // other fields }',properties:{lastTransitionTime:{description:"lastTransitionTime is the last time the condition transitioned from one status to another. This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.",format:"date-time",type:"string"},message:{description:"message is a human readable message indicating details about the transition. This may be an empty string.",maxLength:32768,type:"string"},observedGeneration:{description:"observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance.",format:"int64",minimum:0,type:"integer",maximum:9223372036854776e3},reason:{description:"reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty.",maxLength:1024,minLength:1,pattern:"^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$",type:"string"},status:{description:"status of the condition, one of True, False, Unknown.",enum:["True","False","Unknown"],type:"string"},type:{description:"type of condition in CamelCase or in foo.example.com/CamelCase. --- Many .condition.type values are consistent across resources like Available, but because arbitrary conditions can be useful (see .node.status.conditions), the ability to deconflict is important. The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)",maxLength:316,pattern:"^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$",type:"string"}},required:["lastTransitionTime","message","reason","status","type"],type:"object"},type:"array"},observedGeneration:{format:"int64",type:"integer",minimum:-9223372036854776e3,maximum:9223372036854776e3}},type:"object"}},required:["metadata","spec"],type:"object",$schema:"http://json-schema.org/draft-04/schema#"};S.ITextModel;const Ot=/(config|image|source)s:/,qt=(t,e)=>{if(t.getLineContent(e).search(/\s+resource:/)<0)return null;let s=e;for(;--s>0;){let a=t.getLineContent(s).match(Ot);if(a)return`Cluster${Ae(a[1])}Template`}return null},Wt=(t,e)=>{let r=t.getValue();try{return se(r).spec.resources.filter(a=>a.templateRef.kind===e).map(a=>({insertText:a.name,kind:b.Reference,range:void 0,label:a.name}))}catch{}return[]},Vt=()=>{w.registerCompletionItemProvider("yaml",{triggerCharacters:[" "],provideCompletionItems(t,e){let r=qt(t,e.lineNumber);return r?{incomplete:!1,suggestions:Wt(t,r)}:null}})};window.MonacoEnvironment={getWorker(t,e){switch(e){case"yaml":return new Kt;case"javascript":return new Fe;default:return new Ee}}};Nt({enableSchemaRequest:!0,hover:!0,completion:!0,validate:!0,format:!0,schemas:[{fileMatch:["*.yaml"],schema:Ut}]});Vt();new vt({target:document.getElementById("app")});