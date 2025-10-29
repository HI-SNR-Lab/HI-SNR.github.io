---
title: Research
nav:
  order: 1
  tooltip: Published works
---

# {% include icon.html icon="fa-solid fa-microscope" %}Research

The High-quality Imaging, Signals, N' Radar (HI-SNR) laboratory develops data assimilation, computational imaging, and signal processing methods in two areas:

<details close>
<summary>1. Multistatic Radar for Environmental Applications</summary>
<br>
We develop multistatic radars and algorithms for retreiving properties of the environment for the following purposes:
  1. Equip scientists with the tools needed to improve our understanding of the environment. We develop radar-based methods for making measurements of environmental processes at a quality or scale that is not possible with existing tools. This includes projects such as generating 2D maps of the temperature distribution inside glaciers.
  2. Develop systems than enable resource monitoring and improved resource management. This includes projects such as monitoring soil moisture for improved irrigaiton practices.  

  We are a crib to grave team that puts environmental solutions into the hands of scientists and and environmental agencies. 
</details>

<details close>
<summary>2. Computational Photography</summary>
<br>
We develop computational photography algorithms that improve image quality, make camera performance more equitable, and reconstruct 2-4D images from incomplete data.
</details>



{% include section.html %}
## Research Areas in Computational Photography

{% capture text %}

We develop equitable imaging methods that ensure beautiful, crisp images for all users. From the technical side, this includes improved skin tone mapping, 3A algorithms, and composite photography that adapt to a user's skin tone and the surrounding scene. 


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
  image="images/photography.jpg"
  link="research"
  title="Equitable Imaging"
  flip=true
  style="bare"
  text=text
%}

{% capture text %}

We research human behavioral processes that influence subjective interpretation of image quality and integrate findings from these studies into improved imaging algorithms.


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
  image="images/people_in_photo.jpg"
  link="research"
  title="Behavioral Studies for Improved Imaging"
  flip=true
  style="bare"
  text=text
%}

{% capture text %}

We develop imaging methods and rendering algorithms for reconstructing 3D images from reduced dimensionality data.

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
  image="images/compressive_sensing.png"
  link="research"
  title="Compressive Sensing"
  flip=true
  style="bare"
  text=text
%}


## Research Areas in Environmental Radar

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
  link="research"
  text="See our publications"
  icon="fa-solid fa-arrow-right"
  flip=true
  style="bare"
%}

{% endcapture %}

{%
  include feature.html
  image="images/tree_radar.gif"
  link="research"
  title="Biosphere Radars"
  flip=true
  style="bare"
  text=text
%}

{% capture text %}

We are creating new methods of monitoring our limited resources including drone-based soil moisture estimation and satellite-based aquifer volume retrievals.

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
  image="images/drone_ani.gif"
  link="research"
  title="Near Subsurface Radar"
  text=text
%}



## Highlighted

2020 IGARSS Prize Paper Award Winner  
[Processing-Based Synchronization Approach for Bistatic Radar Glacial Tomography](https://ieeexplore.ieee.org/document/9323969)  
This novel processing based synchronization approach allows coherent long-offset radar processing. This is a key enabling technology for radar tomography of glaciers.   
*N. Bienert, D.. Schroeder, S. Peters, M. Siegfried*

<iframe width="711" height="400" src="https://www.youtube.com/embed/WaF6N42GQiw?si=peiCi045mLVfM2el&amp;controls=0&amp;start=23" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>

{% include section.html %}

## All

{% include search-box.html %}

{% include search-info.html %}

{% include list.html data="citations" component="citation" style="rich" %}
