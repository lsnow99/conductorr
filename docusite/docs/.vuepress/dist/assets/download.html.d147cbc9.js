import{r as t,c as r,a as s,d as e,s as l,F as c,b as n,e as i,o as p}from"./app.2980b5d3.js";import{_ as u}from"./plugin-vue_export-helper.5a098b48.js";const d={},h=s("h1",{id:"download-and-installation",tabindex:"-1"},[s("a",{class:"header-anchor",href:"#download-and-installation","aria-hidden":"true"},"#"),n(" Download and Installation")],-1),m=n("Follow the instructions below depending on your system and preferred installation method. For configuration instructions post-installation, check "),_=n("here"),k=n("."),b=s("h2",{id:"manual-installation",tabindex:"-1"},[s("a",{class:"header-anchor",href:"#manual-installation","aria-hidden":"true"},"#"),n(" Manual Installation")],-1),y=n("Download the latest release for your operating system "),f={href:"https://github.com/lsnow99/conductorr/releases",target:"_blank",rel:"noopener noreferrer"},g=n("here"),v=n("."),w=s("p",null,[n("Put the binary file in a suitable location. Simply double click the binary or run "),s("code",null,"./conductorr"),n(" in your terminal to start Conductorr. On some systems you may need to first run "),s("code",null,"chmod +x conductorr"),n(" in your terminal before launching.")],-1),x=s("p",null,"To launch automatically at startup, please refer to your operating system's instructions.",-1),q=n("Navigate to "),C={href:"http://localhost:6416/",target:"_blank",rel:"noopener noreferrer"},N=n("http://localhost:6416/"),B=n(" and you should see the setup screen prompting you to create the admin user."),D=i(`<h2 id="docker-compose" tabindex="-1"><a class="header-anchor" href="#docker-compose" aria-hidden="true">#</a> Docker Compose</h2><p>Create a <code>docker-compose.yml</code> file with the following contents:</p><div class="language-yaml ext-yml line-numbers-mode"><pre class="language-yaml"><code><span class="token key atrule">version</span><span class="token punctuation">:</span> <span class="token string">&quot;3.9&quot;</span>
<span class="token key atrule">services</span><span class="token punctuation">:</span>
  <span class="token key atrule">conductorr</span><span class="token punctuation">:</span>
    <span class="token key atrule">image</span><span class="token punctuation">:</span> <span class="token string">&quot;conductorr&quot;</span>
    <span class="token key atrule">ports</span><span class="token punctuation">:</span>
      <span class="token punctuation">-</span> <span class="token string">&quot;6416:6416&quot;</span>
    <span class="token key atrule">volumes</span><span class="token punctuation">:</span>
      <span class="token comment"># Replace the . with your preferred location on your system for the database file</span>
      <span class="token punctuation">-</span> .<span class="token punctuation">:</span>/app/conductorr.db  
    <span class="token key atrule">environment</span><span class="token punctuation">:</span>
      <span class="token key atrule">TMDB_API_KEY</span><span class="token punctuation">:</span> yourapikeyhere
      <span class="token comment"># Add any other environment variables for configuration here</span>
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br></div></div>`,3);function F(L,V){const o=t("RouterLink"),a=t("OutboundLink");return p(),r(c,null,[h,s("p",null,[m,e(o,{to:"/guide/configuration.html"},{default:l(()=>[_]),_:1}),k]),b,s("p",null,[y,s("a",f,[g,e(a)]),v]),w,x,s("p",null,[q,s("a",C,[N,e(a)]),B]),D],64)}var R=u(d,[["render",F]]);export{R as default};
