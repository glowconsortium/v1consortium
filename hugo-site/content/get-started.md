---
title: "Get Started with FleetOps"
layout: "simple"
description: "Start your fleet management transformation today. Request access to FleetOps and see how we can streamline your operations."
---

<div class="bg-gradient-to-b from-blue-50 to-white py-16">
  <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
    <div class="text-center mb-12">
      <h1 class="text-4xl font-bold text-gray-900 mb-4">
        Get Started with FleetOps
      </h1>
      <p class="text-xl text-gray-600 max-w-2xl mx-auto">
        Ready to transform your fleet operations? Complete the form below and our team will get you set up with a customized FleetOps solution.
      </p>
    </div>
    <!-- Benefits Section -->
    <div class="grid md:grid-cols-3 gap-6 mb-12">
      <div class="text-center p-6 bg-white rounded-lg shadow-sm">
        <div class="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-4">
          <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
          </svg>
        </div>
        <h3 class="text-lg font-semibold text-gray-900 mb-2">Quick Setup</h3>
        <p class="text-gray-600">Get up and running in 24-48 hours with guided onboarding</p>
      </div>
      <div class="text-center p-6 bg-white rounded-lg shadow-sm">
        <div class="w-12 h-12 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-4">
          <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
        </div>
        <h3 class="text-lg font-semibold text-gray-900 mb-2">Free Trial</h3>
        <p class="text-gray-600">30-day trial with full access to all features and support</p>
      </div>
      <div class="text-center p-6 bg-white rounded-lg shadow-sm">
        <div class="w-12 h-12 bg-purple-100 rounded-full flex items-center justify-center mx-auto mb-4">
          <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
          </svg>
        </div>
        <h3 class="text-lg font-semibold text-gray-900 mb-2">Dedicated Support</h3>
        <p class="text-gray-600">Personal onboarding specialist and 24/7 technical support</p>
      </div>
    </div>
    <!-- Main Form -->
    <div class="bg-white rounded-lg shadow-lg p-8">
      <h2 class="text-2xl font-bold text-gray-900 mb-6">Request Your FleetOps Account</h2> 
      <form action="https://formspree.io/f/placeholder-form-id" method="POST" class="space-y-6">
        <!-- Company Information -->
        <div class="grid md:grid-cols-2 gap-6">
          <div>
            <label for="company-name" class="block text-sm font-medium text-gray-700 mb-1">
              Company Name *
            </label>
            <input 
              type="text" 
              id="company-name" 
              name="company_name" 
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="Enter your company name"
            >
          </div>   
          <div>
            <label for="fleet-size" class="block text-sm font-medium text-gray-700 mb-1">
              Fleet Size *
            </label>
            <select 
              id="fleet-size" 
              name="fleet_size" 
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">Select fleet size</option>
              <option value="1-5">1-5 vehicles</option>
              <option value="6-25">6-25 vehicles</option>
              <option value="26-100">26-100 vehicles</option>
              <option value="101-500">101-500 vehicles</option>
              <option value="500+">500+ vehicles</option>
            </select>
          </div>
        </div>
        <div class="grid md:grid-cols-2 gap-6">
          <div>
            <label for="operation-type" class="block text-sm font-medium text-gray-700 mb-1">
              Operation Type *
            </label>
            <select 
              id="operation-type" 
              name="operation_type" 
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="">Select operation type</option>
              <option value="motor-carrier">Motor Carrier</option>
              <option value="private-fleet">Private Fleet</option>
              <option value="owner-operator">Owner-Operator</option>
              <option value="logistics-company">Logistics Company</option>
              <option value="freight-broker">Freight Broker</option>
              <option value="other">Other</option>
            </select>
          </div>
          <div>
            <label for="mc-number" class="block text-sm font-medium text-gray-700 mb-1">
              MC Number
            </label>
            <input 
              type="text" 
              id="mc-number" 
              name="mc_number"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="MC-123456 (if applicable)"
            >
          </div>
        </div>
        <!-- Contact Information -->
        <h3 class="text-lg font-semibold text-gray-900 mt-8 mb-4">Primary Contact Information</h3>
        <div class="grid md:grid-cols-2 gap-6">
          <div>
            <label for="first-name" class="block text-sm font-medium text-gray-700 mb-1">
              First Name *
            </label>
            <input 
              type="text" 
              id="first-name" 
              name="first_name" 
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="Enter your first name"
            >
          </div>     
          <div>
            <label for="last-name" class="block text-sm font-medium text-gray-700 mb-1">
              Last Name *
            </label>
            <input 
              type="text" 
              id="last-name" 
              name="last_name" 
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="Enter your last name"
            >
          </div>
        </div>
        <div class="grid md:grid-cols-2 gap-6">
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700 mb-1">
              Business Email *
            </label>
            <input 
              type="email" 
              id="email" 
              name="email" 
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="you@yourcompany.com"
            >
          </div>
          <div>
            <label for="phone" class="block text-sm font-medium text-gray-700 mb-1">
              Phone Number *
            </label>
            <input 
              type="tel" 
              id="phone" 
              name="phone" 
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="(555) 123-4567"
            >
          </div>
        </div>
        <div>
          <label for="job-title" class="block text-sm font-medium text-gray-700 mb-1">
            Job Title *
          </label>
          <input 
            type="text" 
            id="job-title" 
            name="job_title" 
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            placeholder="Fleet Manager, Operations Director, etc."
          >
        </div>
        <!-- Business Address -->
        <h3 class="text-lg font-semibold text-gray-900 mt-8 mb-4">Business Address</h3>
        <div>
          <label for="address" class="block text-sm font-medium text-gray-700 mb-1">
            Street Address *
          </label>
          <input 
            type="text" 
            id="address" 
            name="address" 
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            placeholder="123 Main Street"
          >
        </div>
        <div class="grid md:grid-cols-3 gap-6">
          <div>
            <label for="city" class="block text-sm font-medium text-gray-700 mb-1">
              City *
            </label>
            <input 
              type="text" 
              id="city" 
              name="city" 
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="Columbus"
            >
          </div>
          <div>
            <label for="state" class="block text-sm font-medium text-gray-700 mb-1">
              State *
            </label>
            <input 
              type="text" 
              id="state" 
              name="state" 
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="OH"
            >
          </div>
          <div>
            <label for="zip" class="block text-sm font-medium text-gray-700 mb-1">
              ZIP Code *
            </label>
            <input 
              type="text" 
              id="zip" 
              name="zip" 
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="43215"
            >
          </div>
        </div>
        <!-- Additional Information -->
        <h3 class="text-lg font-semibold text-gray-900 mt-8 mb-4">Tell Us About Your Needs</h3>
        <div>
          <label for="current-system" class="block text-sm font-medium text-gray-700 mb-1">
            Current Fleet Management System
          </label>
          <input 
            type="text" 
            id="current-system" 
            name="current_system"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            placeholder="None, Excel, PeopleNet, Omnitracs, etc."
          >
        </div>
        <div>
          <label for="pain-points" class="block text-sm font-medium text-gray-700 mb-1">
            Biggest Challenges
          </label>
          <div class="space-y-2">
            <label class="flex items-center">
              <input type="checkbox" name="challenges" value="compliance" class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
              <span class="ml-2 text-sm text-gray-600">DOT compliance and reporting</span>
            </label>
            <label class="flex items-center">
              <input type="checkbox" name="challenges" value="routing" class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
              <span class="ml-2 text-sm text-gray-600">Route optimization and dispatching</span>
            </label>
            <label class="flex items-center">
              <input type="checkbox" name="challenges" value="communication" class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
              <span class="ml-2 text-sm text-gray-600">Driver communication</span>
            </label>
            <label class="flex items-center">
              <input type="checkbox" name="challenges" value="maintenance" class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
              <span class="ml-2 text-sm text-gray-600">Vehicle maintenance tracking</span>
            </label>
            <label class="flex items-center">
              <input type="checkbox" name="challenges" value="fuel" class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
              <span class="ml-2 text-sm text-gray-600">Fuel cost management</span>
            </label>
            <label class="flex items-center">
              <input type="checkbox" name="challenges" value="paperwork" class="rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
              <span class="ml-2 text-sm text-gray-600">Manual paperwork and processes</span>
            </label>
          </div>
        </div>
        <div>
          <label for="timeline" class="block text-sm font-medium text-gray-700 mb-1">
            Implementation Timeline
          </label>
          <select 
            id="timeline" 
            name="timeline"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="">Select timeline</option>
            <option value="immediate">Immediate (within 30 days)</option>
            <option value="1-3-months">1-3 months</option>
            <option value="3-6-months">3-6 months</option>
            <option value="6-12-months">6-12 months</option>
            <option value="exploring">Just exploring options</option>
          </select>
        </div>
        <div>
          <label for="additional-info" class="block text-sm font-medium text-gray-700 mb-1">
            Additional Information
          </label>
          <textarea 
            id="additional-info" 
            name="additional_info" 
            rows="4"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            placeholder="Tell us more about your specific needs, questions, or anything else we should know..."
          ></textarea>
        </div>
        <!-- Terms and Consent -->
        <div class="space-y-3">
          <label class="flex items-start">
            <input type="checkbox" name="consent_marketing" class="mt-1 rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
            <span class="ml-2 text-sm text-gray-600">
              I would like to receive updates about FleetOps features, industry news, and best practices via email.
            </span>
          </label>      
     <label class="flex items-start">
            <input type="checkbox" name="agree_terms" required class="mt-1 rounded border-gray-300 text-blue-600 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50">
            <span class="ml-2 text-sm text-gray-600">
              I agree to the <a href="/terms" class="text-blue-600 hover:text-blue-800">Terms of Service</a> and <a href="/privacy" class="text-blue-600 hover:text-blue-800">Privacy Policy</a> *
            </span>
          </label>
        </div>
        <!-- Submit Button -->
        <div class="pt-6">
          <button 
            type="submit"
            class="w-full bg-blue-600 text-white font-semibold py-3 px-6 rounded-lg hover:bg-blue-700 transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            Submit Request
          </button>
        </div>
      </form>
    </div>
    <!-- What Happens Next -->
    <div class="mt-12 bg-gray-50 rounded-lg p-8">
      <h3 class="text-xl font-bold text-gray-900 mb-4">What Happens Next?</h3>
      <div class="grid md:grid-cols-3 gap-6">
        <div class="flex items-start">
          <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center mr-3 mt-1">
            <span class="text-blue-600 font-semibold text-sm">1</span>
          </div>
          <div>
            <h4 class="font-semibold text-gray-900 mb-1">Review & Contact</h4>
            <p class="text-gray-600 text-sm">Our team reviews your submission and contacts you within 2 business hours.</p>
          </div>
        </div>  
        <div class="flex items-start">
          <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center mr-3 mt-1">
            <span class="text-blue-600 font-semibold text-sm">2</span>
          </div>
          <div>
            <h4 class="font-semibold text-gray-900 mb-1">Demo & Setup</h4>
            <p class="text-gray-600 text-sm">Schedule a personalized demo and begin account setup with your onboarding specialist.</p>
          </div>
        </div>
        <div class="flex items-start">
          <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center mr-3 mt-1">
            <span class="text-blue-600 font-semibold text-sm">3</span>
          </div>
          <div>
            <h4 class="font-semibold text-gray-900 mb-1">Go Live</h4>
            <p class="text-gray-600 text-sm">Complete onboarding, train your team, and start optimizing your fleet operations.</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<!-- Alternative Contact Section -->
<div class="bg-white py-16">
  <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
    <h2 class="text-2xl font-bold text-gray-900 mb-4">
      Prefer to Talk First?
    </h2>
    <p class="text-gray-600 mb-8">
      Speak directly with our fleet management experts
    </p>
    <div class="space-y-4">
      <div>
        <a href="tel:614-555-0123" class="inline-flex items-center text-blue-600 hover:text-blue-800 font-semibold text-lg">
          <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"></path>
          </svg>
          (614) 555-0123
        </a>
      </div>
      <div>
        <a href="mailto:sales@fleetops.com" class="inline-flex items-center text-blue-600 hover:text-blue-800 font-medium">
          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
          </svg>
          sales@fleetops.com
        </a>
      </div>
    </div>
  </div>
</div>