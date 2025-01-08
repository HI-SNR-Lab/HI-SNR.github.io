---
carousels:
  - images:
    - image: images/students.jpg
    - image: images/ravine.jpg
    - image: images/connecting_apres.jpg
    - image: images/Tents.jpg
    - image: images/testing_with_Neo.jpg
    - image: images/radar_tripod_crop.jpg
---

#  At the High precision Imaging, Signals N' Radar (HI-SNR) Lab, we develop computational imaging techniques that sense our environment at a scale, resolution, or quality that was previously impossible. 

{% include carousel.html height="50" unit="%" duration="5" number="1" %}

{% include section.html %}

## Research Areas

{% capture text %}

We are inventing radar techniques to measure ice properties at scales previously unexplored. The systems she develops provide measurements that better constrain ice sheet models to improve the accuracy of sea level rise predictions.

{%
  include button.html
  link="research"
  text="See our publications"
  icon="fa-solid fa-arrow-right"
  flip=true
  style="bare"
%}

{% endcapture %}

{%
  include feature.html
  image="images/approachingradar.jpg"
  link="research"
  title="Cryosphere Radars"
  text=text
%}

{% capture text %}

We are developing radars and data assimilation techniques for 3D mapping of vegetation structure and water content to improve wildfire risk predictions, CO2 sequestration assessment, and ecosystem health.

{%
  include button.html
  link="projects"
  text="See our publications"
  icon="fa-solid fa-arrow-right"
  flip=true
  style="bare"
%}

{% endcapture %}

{%
  include feature.html
  image="images/tree_radar.gif"
  link="projects"
  title="Biosphere Radars"
  flip=true
  style="bare"
  text=text
%}

{% capture text %}

We are creating new methods of monitoring our limitted resources including drone-based soil moisture estimation and satellite-based aquifer volume retreivals.

{%
  include button.html
  link="team"
  text="See our publications"
  icon="fa-solid fa-arrow-right"
  flip=true
  style="bare"
%}

{% endcapture %}

{%
  include feature.html
  image="images/drone_ani.gif"
  link="team"
  title="Near Subsurface Radar"
  text=text
%}

{% capture text %}

We develop imaging algorithms not only for radar, but also for cameras. We focus on equitable imaging and develop methods that ensure beautiful, crisp images for all users. 


{%
  include button.html
  link="projects"
  text="See our publications"
  icon="fa-solid fa-arrow-right"
  flip=true
  style="bare"
%}

{% endcapture %}

{%
  include feature.html
  image="images/camera_testing.jpg"
  link="projects"
  title="Computational Photography"
  flip=true
  style="bare"
  text=text
%}
