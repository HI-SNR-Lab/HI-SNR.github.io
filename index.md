---
carousels:
  - images:
    - image: images/connecting_apres.jpg
    - image: images/camera_testing.jpg
    - image: images/students.jpg
    - image: images/Yang_Testing.JPG
    - image: images/ravine.jpg
    - image: images/Tents.jpg
    - image: images/holding_camera.jpg
    - image: images/testing_with_Neo.jpg
    - image: images/summer_2025_testing.jpg
---

<style>
.epic-title {
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

.intro-description {
  display: flex;
  flex-direction: column;

  width: 100vw;

  padding: 10px 80px;
  box-sizing: border-box;

  align-items: flex-start;
  text-align: left;
}

.section-orca {
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
</style>

<h1 style="font-size: 40px;"><b>Computational Imaging in Photography and Environmental Radar</b></h1>


<div class="intro-description">
    <h5 style="text-align: center;"><i>Compressed Sensing :small_blue_diamond: Data Fusion :small_blue_diamond: Signal Processing :small_blue_diamond: Adaptive Algorithms :small_blue_diamond: Data Assimilation </i></h5>
</div>

{% include carousel.html height="50" unit="%" duration="5" number="1" %}

{% include section.html %}

<div class="epic-title">
    <h2 style="text-align: center;">At the High precision Imaging, Signals 'N Radar Lab, we develop computational imaging techniques that sense our environment at a scale, resolution, or quality that was previously impossible. </h2>
</div>



## Research Areas
We develop imaging methods, including compressed sensing, data fusion, signal processing, data assimilation, and adaptive algorithms for both camera and radar applications. In both areas, a core focus is image formation with incomplete data.

{% capture text %}

We are inventing radars, signal processing algorithms, and data assimilation techniques to measure the environment at scales previously unexplored. Our systems are deployed to improve geophysical understanding of the earth and to aid in resource management. We have developed systems that inform predictions for sea level rise and wildfire risk, as well as monitoring ecosystem and agricultural health.

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
  title="Environmental Radar"
  text=text
%}

{% capture text %}

We are developing signal processing algorithms, data fusion methods, compressed sensing, and adaptive algorithms to form beautiful photos and renderings that represent our lives. Our algorithms have been used in equitable imaging algorithms that ensure beautiful, crisp images for all users. We are developing compressive sensing methods for reduced cost 3D imaging for medical applications.


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
  image="images/camera_testing.jpg"
  link="research"
  title="Computational Photography"
  flip=true
  style="bare"
  text=text
%}


{% include section.html %}

<div class="epic-title">
{%
  include button.html
  link="orca"
  text=" <h1>Documentation on Open Code Radar Architecture (ORCA)</h1>"
  style="bare"
  background="#ddddee"
%}
</div>