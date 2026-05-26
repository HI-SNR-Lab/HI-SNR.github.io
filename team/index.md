---
title: Team
nav:
  order: 3
  tooltip: About our team
---

<style>
.team-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  align-items: flex-start;
}

h1, h2 {
  margin-top: 4px !important;
  margin-bottom: 4px !important;
}

.section-divider {
  margin: 4px 0 !important;
  padding: 4;
}

.section-pi {
  display: flex;
  margin-left: 4px 0 !important;
  flex-direction: column;
  align-items: center;
  text-align: center;
  width: 100vw;
  margin-left: calc(-50vw + 50%);
  padding: 0px 0;
  background: #eef2ff
}

.section-grad {
  display: flex;
  flex-direction: column;

  width: 100vw;
  margin-left: calc(50% - 50vw);

  padding: 20px 80px;
  box-sizing: border-box;

  align-items: flex-start;
  text-align: left;

  background: #ddddee;
}

.section-undergrad {
  display: flex;
  flex-direction: column;

  width: 100vw;
  margin-left: calc(50% - 50vw);

  padding: 20px 80px;
  box-sizing: border-box;

  align-items: flex-start;
  text-align: left;

  background: #eef1ff;
}

.section-alumni {
  display: flex;
  flex-direction: column;

  width: 100vw;
  margin-left: calc(50% - 50vw);

  padding: 20px 80px;
  box-sizing: border-box;

  align-items: flex-start;
  text-align: left;

  background: #ddddee;
}

.section-past {
  display: flex;
  flex-direction: column;

  width: 100vw;
  margin-left: calc(50% - 50vw);

  padding: 20px 80px;
  box-sizing: border-box;

  align-items: flex-start;
  text-align: left;

  background: #eef2ff;
}

/* .section {
  width: 100vw;
  margin-left: calc(-50vw + 50%);
  padding: 30px 0;
} */
</style>


# {% include icon.html icon="fa-solid fa-users" %}Team

At HI-SNR Lab, we believe that a joyful and inclusive environment drives innovation. Our team is dedicated to collaboration and creativity, celebrating our differences, and supporting one another. Join us on our journey to bring new imaging systems and methods to our communities!

{% include section.html %}

<div class="section section-pi"> 
  <div class="inner"> 
    <h2 class="section-title">Principal Investigator</h2>
    <div class="team-grid">
      {% for member in site.members %} {% if member.role == "pi" %}
        <div class="team-member"> {% include portrait.html lookup=member.slug %} </div>
      {% endif %} {% endfor %}  
    </div>
  </div> 
</div>

<div class="section section-grad"> 
  <div class="inner"> 
    <h2 class="section-title">Graduate Students</h2>
    <div class="team-grid">
      {% for member in site.members %} {% if member.role == "grad" %}
        <div class="team-member"> {% include portrait.html lookup=member.slug %} </div>
      {% endif %} {% endfor %}  
    </div>
  </div> 
</div>

<div class="section section-undergrad"> 
  <div class="inner"> 
    <h2 class="section-title">Undergraduate Students</h2>
    <div class="team-grid">
      {% for member in site.members %} {% if member.role == "undergrad" %}
        <div class="team-member"> {% include portrait.html lookup=member.slug %} </div>
      {% endif %} {% endfor %}  
    </div>
  </div> 
</div>

<div class="section section-alumni"> 
  <div class="inner"> 
    <h2 class="section-title">Alumni</h2>
    <div class="team-grid">
      {% for member in site.members %} {% if member.role == "alumni" %}
        <div class="team-member"> {% include portrait.html lookup=member.slug %} </div>
      {% endif %} {% endfor %}  
    </div>
  </div> 
</div>

<div class="section section-past"> 
  <div class="inner"> 
    <h2 class="section-title">Past Students</h2>
    <div class="team-grid">
      {% for member in site.members %} {% if member.role == "past" %}
        <div class="team-member"> {% include portrait.html lookup=member.slug %} </div>
      {% endif %} {% endfor %}  
    </div>
  </div> 
</div>

